package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/rest"
)

func main() {

	cf := genericclioptions.NewConfigFlags(true)

	// disable warnings
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
	fmt.Println("---------------------------------------------")
	fmt.Println(fmt.Sprintf("ServerVersion:%s", serverVersion))

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	etcdPods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{LabelSelector: "component=etcd"})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range etcdPods.Items {
		etcdImage := pod.Spec.Containers[0].Image
		fmt.Println(fmt.Sprintf("etcd Name:%s \n ContainerName:%s", pod.Spec.Containers[0].Name, etcdImage))
	}

	cmPods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{LabelSelector: "component=kube-controller-manager"})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range cmPods.Items {
		cmImage := pod.Spec.Containers[0].Image
		fmt.Println(fmt.Sprintf("controller-manager Name:%s \n ContainerName:%s", pod.Spec.Containers[0].Name, cmImage))
	}

	schedulerPod, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{LabelSelector: "component=kube-scheduler"})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range schedulerPod.Items {
		schedulerImage := pod.Spec.Containers[0].Image
		fmt.Println(fmt.Sprintf("scheduler Name:%s \n ContainerName:%s", pod.Spec.Containers[0].Name, schedulerImage))
	}

	fmt.Println("---------------------------------------------")
	for _, node := range nodes.Items {
		name := node.Name
		kubeletVersion := node.Status.NodeInfo.KubeletVersion
		containerRuntimeVersion := node.Status.NodeInfo.ContainerRuntimeVersion
		osImage := node.Status.NodeInfo.OSImage
		hostname := node.ObjectMeta.Name
		kernelVersion := node.Status.NodeInfo.KernelVersion
		kubeProxyVersion := node.Status.NodeInfo.KubeProxyVersion
		operatingSystem := node.Status.NodeInfo.OperatingSystem
		architecture := node.Status.NodeInfo.Architecture

		fmt.Println("---------------------------------------------")
		fmt.Println(fmt.Sprintf("Name:%s", name))
		fmt.Println(fmt.Sprintf("containerRuntimeVersion:%s", containerRuntimeVersion))
		fmt.Println(fmt.Sprintf("osImage:%s", osImage))
		fmt.Println(fmt.Sprintf("hostname:%s", hostname))
		fmt.Println(fmt.Sprintf("kubeletVersion:%s", kubeletVersion))
		fmt.Println(fmt.Sprintf("kernelVersion:%s", kernelVersion))
		fmt.Println(fmt.Sprintf("kubeProxyVersion:%s", kubeProxyVersion))
		fmt.Println(fmt.Sprintf("operatingSystem:%s", operatingSystem))
		fmt.Println(fmt.Sprintf("architecture:%s", architecture))

		// Do something with the version information
	}

	addons := []string{
		"kube-dns",
		"metrics-server",
		"dashboard",
		// Add other add-on names as needed
	}
	fmt.Println("-------------------- k8s addons -------------------------------")
	for _, addon := range addons {
		addonPods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{LabelSelector: "k8s-app=" + addon})
		if err != nil {
			panic(err.Error())
		}
		for _, pod := range addonPods.Items {
			addonImage := pod.Spec.Containers[0].Image
			fmt.Println(fmt.Sprintf("Pod Name:%s \n ContainerName:%s", pod.Name, addonImage))
			// Do something with the version information
		}
	}
}
