package main

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	estorePdt "github.com/arutselvan15/estore-product-kube-client/pkg/client/clientset/versioned"
)

func main() {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"

	// Use kubeconfig to create client config.
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pdtClientset, err := estorePdt.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// list namespaces
	nsList, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, ns := range nsList.Items {
		fmt.Println(ns.Name)
	}

	// list eShop product resources
	pdtList, err := pdtClientset.EstoreV1().Products("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, ns := range pdtList.Items {
		fmt.Println(ns.Namespace, ns.Name)
	}
}
