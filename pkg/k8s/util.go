package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/google/go-containerregistry/pkg/name"
	containerimage "github.com/google/go-containerregistry/pkg/name"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
)

const (
	k8sComponentNamespace = "kube-system"
)

func CollectNodes(clientset *kubernetes.Clientset) []NodeInfo {
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	nodesInfo := make([]NodeInfo, 0)
	for _, node := range nodes.Items {
		nodeRole := "worker"
		if _, ok := node.Labels["node-role.kubernetes.io/control-plane"]; ok {
			nodeRole = "master"
		}
		if _, ok := node.Labels["node-role.kubernetes.io/master"]; ok {
			nodeRole = "master"
		}
		nodesInfo = append(nodesInfo, NodeInfo{
			NodeName:                node.Name,
			KubeletVersion:          node.Status.NodeInfo.KubeletVersion,
			ContainerRuntimeVersion: node.Status.NodeInfo.ContainerRuntimeVersion,
			OsImage:                 node.Status.NodeInfo.OSImage,
			Hostname:                node.ObjectMeta.Name,
			KernelVersion:           node.Status.NodeInfo.KernelVersion,
			KubeProxyVersion:        node.Status.NodeInfo.KernelVersion,
			OperatingSystem:         node.Status.NodeInfo.OperatingSystem,
			Architecture:            node.Status.NodeInfo.Architecture,
			NodeRole:                nodeRole,
		})
	}
	return nodesInfo
}

func CollectCoreComponents[T any](components []T, clientset *kubernetes.Clientset, getComponent func(name.Reference, name.Reference) (T, error)) ([]T, error) {
	labelSelector := "component"
	pods := GetPodsInfo(clientset, labelSelector, k8sComponentNamespace)

	for _, pod := range pods.Items {
		for _, s := range pod.Status.ContainerStatuses {
			imageRef, err := containerimage.ParseReference(s.ImageID)
			if err != nil {
				return nil, err
			}
			imageName, err := containerimage.ParseReference(s.Image)
			if err != nil {
				return nil, err
			}
			c, err := getComponent(imageRef, imageName)
			if err != nil {
				continue
			}
			components = append(components, c)
		}

	}
	return components, nil
}

func GetPodsInfo(clientset *kubernetes.Clientset, labelSelector string, namespace string) *corev1.PodList {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		panic(err.Error())
	}
	return pods
}

func CollectAddons[T any](addons []T, clientset *kubernetes.Clientset, getComponent func(name.Reference, name.Reference) (T, error)) ([]T, error) {
	labelSelector := "k8s-app"
	pods := GetPodsInfo(clientset, labelSelector, k8sComponentNamespace)
	for _, pod := range pods.Items {
		for _, s := range pod.Status.ContainerStatuses {
			imageRef, err := containerimage.ParseReference(s.ImageID)
			if err != nil {
				return nil, err
			}
			imageName, err := containerimage.ParseReference(s.Image)
			if err != nil {
				return nil, err
			}
			c, err := getComponent(imageRef, imageName)
			if err != nil {
				continue
			}
			addons = append(addons, c)
		}
	}
	return addons, nil
}

func CollectOpenShiftComponents[T any](components []T, clientset *kubernetes.Clientset, getComponent func(name.Reference, name.Reference) (T, error)) ([]T, error) {
	namespaceLabel := map[string]string{
		"openshift-kube-apiserver":          "apiserver",
		"openshift-kube-controller-manager": "kube-controller-manager",
		"openshift-kube-scheduler":          "scheduler",
		"openshift-etcd":                    "etcd",
	}
	for namespace, labelSelector := range namespaceLabel {
		pods := GetPodsInfo(clientset, labelSelector, namespace)

		for _, pod := range pods.Items {
			for _, s := range pod.Status.ContainerStatuses {
				imageRef, err := containerimage.ParseReference(s.ImageID)
				if err != nil {
					return nil, err
				}
				imageName, err := containerimage.ParseReference(s.Image)
				if err != nil {
					return nil, err
				}
				c, err := getComponent(imageRef, imageName)
				if err != nil {
					continue
				}
				components = append(components, c)
			}

		}
	}
	return components, nil
}

func CurrentTimeStamp() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}

func GetMetadata[T any](clusterName string, serverVersion *version.Info, getMetadata func(string, *version.Info) T) T {
	return getMetadata(clusterName, serverVersion)
}

func BomToCycloneDx(cluster *Cluster) *cdx.BOM {
	cMetadeata := cluster.Metadata
	tools := make([]cdx.Tool, 0)
	for _, t := range cMetadeata.Tools {
		tools = append(tools, cdx.Tool{
			Vendor:  t.Vendor,
			Name:    t.Name,
			Version: t.Version,
		})
	}
	metadata := cdx.Metadata{
		Timestamp: cMetadeata.Timestamp,
		Tools:     &tools,
		Component: &cdx.Component{
			BOMRef:  fmt.Sprintf("pkg:%s:%s", cMetadeata.Component.Name, strings.Replace(cMetadeata.Component.Version, "v", "", -1)),
			Type:    cdx.ComponentTypeApplication,
			Name:    cMetadeata.Component.Name,
			Version: cMetadeata.Component.Version,
		},
	}
	cdxComponents := make([]cdx.Component, 0)
	allComponents := make([]Component, 0)
	allComponents = append(allComponents, cluster.ControlPlane.Components...)
	allComponents = append(allComponents, cluster.Addons...)
	for _, c := range allComponents {
		bomRef := fmt.Sprintf("pkg:oci/%s@%s?repository_url=%s/library/%s",
			c.Repository,
			c.Digest,
			c.Registry,
			c.Repository)
		cdxComponents = append(cdxComponents, cdx.Component{
			BOMRef:     bomRef,
			Type:       cdx.ComponentTypeContainer,
			Name:       c.Repository,
			Version:    c.Version,
			PackageURL: bomRef,
		})
	}
	cdxDependecies := GetSbomDependency(cluster.Metadata.Component.Name, cdxComponents)
	nodeComponents, nodeDependecies := GettNodeComponentAndDependency(cluster.NodesInfo)
	cdxDependecies = append(cdxDependecies, nodeDependecies...)
	cdxComponents = append(cdxComponents, nodeComponents...)
	return CreateCycloneDXSbom(metadata, cdxDependecies, cdxComponents)
}

func GetSbomDependency(ref string, components []cdx.Component) []cdx.Dependency {
	dependencies := make([]cdx.Dependency, 0)
	dependsOn := make([]string, 0)
	for _, c := range components {
		dependsOn = append(dependsOn, c.BOMRef)
	}
	dependencies = append(dependencies, cdx.Dependency{
		Ref:          ref,
		Dependencies: &dependsOn,
	})
	return dependencies
}

func GettNodeComponentAndDependency(nodesInfo []NodeInfo) ([]cdx.Component, []cdx.Dependency) {
	components := make([]cdx.Component, 0)
	dependencies := make([]cdx.Dependency, 0)
	for _, n := range nodesInfo {
		nodePurl := fmt.Sprintf("pkg:%s", n.NodeName)
		nodeComponent := cdx.Component{
			Name:       n.NodeName,
			BOMRef:     nodePurl,
			PackageURL: nodePurl,
			Type:       cdx.ComponentTypeContainer,
			Properties: &[]cdx.Property{
				{
					Name:  "node-role",
					Value: n.NodeRole,
				},
				{
					Name:  "host_name",
					Value: n.Hostname,
				},
			},
		}
		components = append(components, nodeComponent)
		osParts := strings.Split(n.OsImage, " ")
		if len(osParts) > 1 {
			purl := fmt.Sprintf("pkg:%s@v%s", osParts[0], osParts[1])
			osComponent := cdx.Component{
				BOMRef:     purl,
				Name:       osParts[0],
				Type:       cdx.ComponentTypeOS,
				Version:    osParts[1],
				PackageURL: purl,
				Properties: &[]cdx.Property{
					{
						Name:  "architecture",
						Value: n.Architecture,
					},
					{
						Name:  "kernel_version",
						Value: n.KernelVersion,
					},
					{
						Name:  "operating_system",
						Value: n.OperatingSystem,
					},
				},
			}
			components = append(components, osComponent)
			dependencies = append(dependencies, cdx.Dependency{
				Ref:          nodeComponent.PackageURL,
				Dependencies: &[]string{osComponent.BOMRef},
			})

			kubletPurl := fmt.Sprintf("pkg:kubelet@%s", n.KubeletVersion)
			kubletComponent := cdx.Component{
				BOMRef:     kubletPurl,
				Name:       "kubelet",
				Type:       cdx.ComponentTypeLibrary,
				Version:    n.KernelVersion,
				PackageURL: kubletPurl,
			}
			dependencies = append(dependencies, cdx.Dependency{
				Ref:          osComponent.BOMRef,
				Dependencies: &[]string{kubletComponent.BOMRef},
			})
			components = append(components, kubletComponent)
			kubeProxyPurl := fmt.Sprintf("pkg:kube-proxy@v%s", n.KubeProxyVersion)
			kubeProxyComponent := cdx.Component{
				BOMRef:     kubeProxyPurl,
				Name:       "kube-proxy",
				Type:       cdx.ComponentTypeLibrary,
				Version:    n.KubeProxyVersion,
				PackageURL: kubeProxyPurl,
			}
			dependencies = append(dependencies, cdx.Dependency{
				Ref:          osComponent.BOMRef,
				Dependencies: &[]string{kubeProxyComponent.BOMRef},
			})
			components = append(components, kubeProxyComponent)

			containerdParts := strings.Split(n.ContainerRuntimeVersion, "://")
			continerdPurl := fmt.Sprintf("pkg:%s@v%s", containerdParts[0], containerdParts[1])
			containerdComponent := cdx.Component{
				BOMRef:     continerdPurl,
				Name:       containerdParts[0],
				Version:    containerdParts[1],
				Type:       cdx.ComponentTypeLibrary,
				PackageURL: continerdPurl,
			}
			if len(containerdParts) > 1 {
				components = append(components, containerdComponent)
			}
			dependencies = append(dependencies, cdx.Dependency{
				Ref:          osComponent.BOMRef,
				Dependencies: &[]string{containerdComponent.BOMRef},
			})
		}
	}
	return components, dependencies
}

func CreateCycloneDXSbom(metadata cdx.Metadata, dependencies []cdx.Dependency, components []cdx.Component) *cdx.BOM {
	bom := cdx.NewBOM()
	bom.Metadata = &metadata
	bom.Components = &components
	bom.Dependencies = &dependencies
	return bom
}

func WriteOutput(clsuter *Cluster, bomType string, format string) error {
	if bomType == "cyclonedx" {
		cdxBom := BomToCycloneDx(clsuter)

		bomFormat := cdx.BOMFileFormatJSON
		if format == "xml" {
			bomFormat = cdx.BOMFileFormatXML
		}
		return cdx.NewBOMEncoder(os.Stdout, bomFormat).
			SetPretty(true).
			Encode(cdxBom)
	}
	b, err := json.MarshalIndent(&clsuter, "", "	")
	if err != nil {
		return err
	}
	fmt.Print(string(b))
	return nil
}
