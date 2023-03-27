package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
	depComponents := make([]*k8s.Component, 0)
	depComponents = append(depComponents, components...)
	depComponents = append(depComponents, addons...)
	clusterName := rawCfg.Contexts[rawCfg.CurrentContext].Cluster
	dependencies := k8s.GetDependencies(clusterName, depComponents)
	now := time.Now()
	ftime := now.Format(time.RFC3339)
	k8sCluster := &k8s.Cluster{
		BomFormat:    "CycloneDX",
		SpecVersion:  "1.4",
		SerialNumber: "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79",
		Version:      1,
		Metadata: k8s.Metadata{
			Timestamp: ftime,
			Tools: []k8s.Tool{
				{
					Vendor:  "aquasecurity",
					Name:    "trivy",
					Version: "0.38.1",
				},
			},
			Component: k8s.Component{
				BomRef:  fmt.Sprintf("kubernetes:%s", strings.Replace(serverVersion.GitVersion, "v", "", -1)),
				Name:    clusterName,
				Type:    "Cluster",
				Version: serverVersion.GitVersion,
			},
		},
		ControlPlane: k8s.ControlPlane{Components: components},
		NodesInfo:    nodesInfo,
		Addons:       addons,
		Dependencies: dependencies,
	}
	b, err := json.Marshal(k8sCluster)
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(string(b))
}
