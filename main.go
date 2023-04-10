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

	args := os.Args
	// collect nodes info
	c := k8s.NewCluster(clientset, clientConfig)
	clusterType := ""
	if len(args) > 2 && args[2] == "ocp" {
		clusterType = "ocp"
	}
	clusterBom, err := c.CreateClusterSbom(clusterType)
	if err != nil {
		panic(err.Error())
	}
	bomType := ""
	if len(args) > 1 {
		bomType = args[1]
	}
	err = k8s.WriteOutput(clusterBom, bomType, "json")
	if err != nil {
		panic(err.Error())
	}

}
