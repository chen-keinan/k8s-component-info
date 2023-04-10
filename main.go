package main

import (
	"os"

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
	bom := k8s.CreateBasicBom(serverVersion, metadata, components, addons, nodesInfo)
	bomType := ""
	if len(args) > 1 {
		bomType = args[1]
	}
	err = k8s.WriteOutput(bom, bomType, "json")
	if err != nil {
		panic(err.Error())
	}

}
