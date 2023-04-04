package main

import (
	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/chen-keinan/k8s-component-info/pkg/k8s"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
)

func main() {
	cf := genericclioptions.NewConfigFlags(true)
	rest.SetDefaultWarningHandler(rest.NoWarnings{})
	clientConfig := cf.ToRawKubeConfigLoader()
	rc, err := clientConfig.ClientConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(rc)
	if err != nil {
		panic(err.Error())
	}

	serverVersion, err := clientset.ServerVersion()
	if err != nil {
		panic(err.Error())
	}
	var components []cdx.Component
	args := os.Args
	if len(args) > 1 && args[1] == "ocp" {
		components = k8s.CollectOpenShiftComponents(clientset)
	} else {
		// collect core components
		components = k8s.CollectCoreComponents(clientset)
	}
	// collect nodes info
	nodesInfo := k8s.CollectNodes(clientset)
	// collect addons info
	addons := k8s.CollectAddons(clientset)
	rawCfg, err := clientConfig.RawConfig()
	if err != nil {
		panic(err.Error())
	}
	depComponents := make([]cdx.Component, 0)
	depComponents = append(depComponents, components...)
	depComponents = append(depComponents, addons...)
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	metadata := k8s.GetSbomMetadata(clusterName, serverVersion)
	dependencies := k8s.GetSbomDependency(metadata.Component.BOMRef, depComponents)
	nodeComponent, nodeDepndencies := k8s.GettNodeComponentAndDependency(nodesInfo)
	dependencies = append(dependencies, nodeDepndencies...)
	depComponents = append(depComponents, nodeComponent...)
	bom := cdx.NewBOM()
	bom.Metadata = &metadata
	bom.Components = &depComponents
	//bom.Components = &components
	bom.Dependencies = &dependencies

	err = cdx.NewBOMEncoder(os.Stdout, cdx.BOMFileFormatJSON).
		SetPretty(true).
		Encode(bom)
	if err != nil {
		panic(err)
	}
}
