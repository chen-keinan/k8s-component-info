package k8s

import (
	"context"
	"strings"
	"time"

	"fmt"

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

func GetPodsInfo(clientset *kubernetes.Clientset, labelSelector string, namespace string) *corev1.PodList {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		panic(err.Error())
	}
	return pods
}

func GetDependencies(ref string, components []cdx.Component) []Dependency {
	dependencies := make([]Dependency, 0)
	dependsOn := make([]string, 0)
	for _, c := range components {
		dependsOn = append(dependsOn, fmt.Sprintf("pkg:%s/%s", c.Type, fmt.Sprintf("%s@%s", c.Name, c.Version)))
	}
	dependencies = append(dependencies, Dependency{
		Ref:       ref,
		DependsOn: dependsOn,
	})
	return dependencies
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

func CreateSbom() {
	metadata := cdx.Metadata{
		// Define metadata about the main component
		// (the component which the BOM will describe)
		Component: &cdx.Component{
			BOMRef:  "pkg:golang/acme-inc/acme-app@v1.0.0",
			Type:    cdx.ComponentTypeApplication,
			Name:    "ACME Application",
			Version: "v1.0.0",
		},
		// Use properties to include an internal identifier for this BOM
		// https://cyclonedx.org/use-cases/#properties--name-value-store
		Properties: &[]cdx.Property{
			{
				Name:  "internal:bom-identifier",
				Value: "123456789",
			},
		},
	}

	// Define the components that acme-app ships with
	// https://cyclonedx.org/use-cases/#inventory

	// Define the dependency graph
	// https://cyclonedx.org/use-cases/#dependency-graph
	dependencies := []cdx.Dependency{
		{
			Ref: "pkg:golang/acme-inc/acme-app@v1.0.0",
			Dependencies: &[]string{
				"pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
			},
		},
		{
			Ref: "pkg:golang/github.com/CycloneDX/cyclonedx-go@v0.3.0",
		},
	}

	// Assemble the BOM
	bom := cdx.NewBOM()
	bom.Metadata = &metadata
	//bom.Components = &components
	bom.Dependencies = &dependencies
}

func GetSbomMetadata(clusterName string, serverVersion *version.Info) cdx.Metadata {
	now := time.Now()
	ftime := now.Format(time.RFC3339)
	return cdx.Metadata{
		Timestamp: ftime,
		Component: &cdx.Component{
			BOMRef:  fmt.Sprintf("pkg:%s:%s", clusterName, strings.Replace(serverVersion.GitVersion, "v", "", -1)),
			Type:    cdx.ComponentTypeApplication,
			Name:    clusterName,
			Version: serverVersion.GitVersion,
		},
	}
}

func GetSbomComponent(imageRef name.Reference, imageName name.Reference) (cdx.Component, error) {
	repoName := imageRef.Context().RepositoryStr()
	registryName := imageRef.Context().RegistryStr()
	if strings.HasPrefix(repoName, "library/sha256") {
		repoName = imageName.Context().RepositoryStr()
		registryName = imageName.Context().RegistryStr()
	}
	bomRef := fmt.Sprintf("pkg:oci/%s@%s?repository_url=%s/library/%s",
		repoName,
		imageRef.Context().Digest(imageRef.Identifier()).DigestStr(),
		registryName,
		repoName)
	return cdx.Component{
		BOMRef:     bomRef,
		Type:       cdx.ComponentTypeContainer,
		Name:       repoName,
		Version:    imageName.Identifier(),
		PackageURL: bomRef,
	}, nil
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
