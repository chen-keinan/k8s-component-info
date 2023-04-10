package main

import (
	"fmt"
	"os"

	"github.com/chen-keinan/k8s-component-info/pkg/k8s"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/spf13/cobra"
)

func main() {
	Execute()
}

var report string
var format string

func init() {
	rootCmd.Flags().StringVarP(&report, "report", "r", "", "--report cyclonedx")
	rootCmd.Flags().StringVarP(&format, "format", "f", "json", "--format json | xml")
}

var rootCmd = &cobra.Command{
	Use:   "kbom",
	Short: "kbom - a simple CLI to produce k8s bill of materials",
	Long:  "kbom - collect k8s core component , addon and nodes info and produce bil of materials",
	RunE: func(cmd *cobra.Command, args []string) error {
		cf := genericclioptions.NewConfigFlags(true)
		rest.SetDefaultWarningHandler(rest.NoWarnings{})
		clientConfig := cf.ToRawKubeConfigLoader()
		rc, err := clientConfig.ClientConfig()
		if err != nil {
			return err
		}
		clientset, err := kubernetes.NewForConfig(rc)
		if err != nil {
			return err
		}
		// collect nodes info
		c := k8s.NewCluster(clientset, clientConfig)
		clusterBom, err := c.CreateClusterSbom()
		if err != nil {
			panic(err.Error())
		}
		err = k8s.WriteOutput(clusterBom, report, format)
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
