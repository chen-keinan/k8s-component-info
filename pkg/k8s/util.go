package k8s

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	cdx "github.com/CycloneDX/cyclonedx-go"
)

func CurrentTimeStamp() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}

func BomToCycloneDx(cluster *ClusterBom) *cdx.BOM {
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

func WriteOutput(clsuter *ClusterBom, bomType string, format string) error {
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
