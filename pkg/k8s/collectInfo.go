package k8s

import (
	"fmt"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	"k8s.io/apimachinery/pkg/version"
)

func GetBaseComponent(imageRef name.Reference, imageName name.Reference) (Component, error) {
	repoName := imageRef.Context().RepositoryStr()
	registryName := imageRef.Context().RegistryStr()
	if strings.HasPrefix(repoName, "library/sha256") {
		repoName = imageName.Context().RepositoryStr()
		registryName = imageName.Context().RegistryStr()
	}

	return Component{
		Repository: repoName,
		Registry:   registryName,
		Name:       fmt.Sprintf("%s:%s", repoName, imageName.Identifier()),
		Digest:     imageRef.Context().Digest(imageRef.Identifier()).DigestStr(),
		Version:    imageName.Identifier(),
	}, nil
}

func GetBasicMetadata(clusterName string, serverVersion *version.Info) Metadata {
	tools := []Tool{
		{
			Vendor:  "aquasecurity",
			Name:    "trivy",
			Version: "0.38.1",
		},
	}
	return Metadata{
		Timestamp: CurrentTimeStamp(),
		Tools:     tools,
		Component: Component{
			Name:    clusterName,
			Version: serverVersion.GitVersion,
		},
	}
}

func CreateBasicBom(version *version.Info, metadata Metadata, components []Component, Addons []Component, nodeInfo []NodeInfo) *Cluster {
	bom := &Cluster{
		Metadata: metadata,
		ControlPlane: ControlPlane{
			Components: components,
		},
		NodesInfo: nodeInfo,
		Addons:    Addons,
	}
	return bom
}
