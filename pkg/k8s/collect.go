package k8s

import (
	"strings"

	"fmt"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/google/go-containerregistry/pkg/name"

	"k8s.io/apimachinery/pkg/version"
)

func GetSbomMetadata(clusterName string, serverVersion *version.Info) cdx.Metadata {
	tools := &[]cdx.Tool{
		{
			Vendor:  "aquasecurity",
			Name:    "trivy",
			Version: "0.38.1",
		},
	}
	return cdx.Metadata{
		Timestamp: CurrentTimeStamp(),
		Tools:     tools,
		Component: &cdx.Component{
			BOMRef: fmt.Sprintf("pkg:%s:%s", clusterName, strings.Replace(serverVersion.GitVersion, "v", "", -1)),
			Type:   cdx.ComponentTypeApplication,
			Name:   clusterName,

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
