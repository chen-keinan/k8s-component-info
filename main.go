package main

import (
	"encoding/json"
	"fmt"

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
	// collect core components
	components := k8s.CollectCoreComponents(clientset)
	// collect nodes info
	nodesInfo := k8s.CollectNodes(clientset)
	// collect addons info
	addons := k8s.CollectAddons(clientset)

	rawCfg, err := clientConfig.RawConfig()
	if err != nil {
		panic(err.Error())
	}
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	k8sCluster := &k8s.Cluster{
		ClusterName:  clusterName,
		Version:      serverVersion,
		ControlPlane: k8s.ControlPlane{Components: components},
		NodesInfo:    nodesInfo,
		Addons:       addons,
	}
	b, err := json.Marshal(k8sCluster)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(string(b))
}
