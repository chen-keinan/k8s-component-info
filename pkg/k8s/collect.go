package k8s

import (
	"strings"
	"time"

	"fmt"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/google/go-containerregistry/pkg/name"

	"k8s.io/apimachinery/pkg/version"
)

const (
	k8sComponentNamespace = "kube-system"
)

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

func CreateCycloneDXSbom(metadata cdx.Metadata, dependencies []cdx.Dependency, components []cdx.Component) *cdx.BOM {
	// Assemble the BOM
	bom := cdx.NewBOM()
	bom.Metadata = &metadata
	bom.Components = &components
	//bom.Components = &components
	bom.Dependencies = &dependencies
	return bom
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
