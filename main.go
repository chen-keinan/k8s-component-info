package main

import (
	"encoding/json"
	"fmt"
	"os"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/chen-keinan/k8s-component-info/pkg/k8s"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
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

	args := os.Args
	// collect nodes info
	nodesInfo := k8s.CollectNodes(clientset)
	// collect addons info

	rawCfg, err := clientConfig.RawConfig()
	if err != nil {
		panic(err.Error())
	}
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	if args[1] == "cyclonedx" {
		var components []cdx.Component
		cdxComponents := make([]cdx.Component, 0)
		if len(args) > 2 && len(args) > 1 && args[2] == "ocp" {
			components, err = k8s.CollectOpenShiftComponents(cdxComponents, clientset, k8s.GetSbomComponent)
			if err != nil {
				panic(err.Error())
			}
		} else {
			// collect core components
			components, err = k8s.CollectCoreComponents(cdxComponents, clientset, k8s.GetSbomComponent)
			if err != nil {
				panic(err.Error())
			}
		}
		addons, err := k8s.CollectAddons(cdxComponents, clientset, k8s.GetSbomComponent)
		if err != nil {
			panic(err.Error())
		}
		components = append(components, addons...)
		metadata := k8s.GetMetadata(clusterName, serverVersion, k8s.GetSbomMetadata)
		dependencies := k8s.GetSbomDependency(metadata.Component.BOMRef, components)
		nodeComponent, nodeDepndencies := k8s.GettNodeComponentAndDependency(nodesInfo)
		components = append(components, nodeComponent...)
		dependencies = append(dependencies, nodeDepndencies...)
		bom := k8s.CreateCycloneDXSbom(metadata, dependencies, components)
		err = cdx.NewBOMEncoder(os.Stdout, cdx.BOMFileFormatJSON).
			SetPretty(true).
			Encode(bom)
		if err != nil {
			panic(err)
		}
	} else {
		baseComponents := make([]k8s.Component, 0)
		addonComponents := make([]k8s.Component, 0)
		var components []k8s.Component
		if len(args) > 2 && len(args) > 1 && args[2] == "ocp" {
			components, err = k8s.CollectOpenShiftComponents(baseComponents, clientset, k8s.GetBaseComponent)
			if err != nil {
				panic(err.Error())
			}
		} else {
			// collect core components
			components, err = k8s.CollectCoreComponents(baseComponents, clientset, k8s.GetBaseComponent)
			if err != nil {
				panic(err.Error())
			}
		}
		metadata := k8s.GetMetadata(clusterName, serverVersion, k8s.GetBasicMetadata)
		addons, err := k8s.CollectAddons(addonComponents, clientset, k8s.GetBaseComponent)
		if err != nil {
			panic(err.Error())
		}
		bom := k8s.CreateBasicSbom(serverVersion, metadata, components, addons, nodesInfo)
		b, err := json.Marshal(&bom)
		fmt.Print(string(b))
	}
}
