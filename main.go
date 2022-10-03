package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/vikash/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		log.Print(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Print(err)
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}

	deploy, err := clientset.AppsV1().Deployments("kube-system").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Print(err)
	}

	//get deployments
	fmt.Println("deployments from default namespaces")
	for _, deploy := range deploy.Items {
		fmt.Printf("%s\n", deploy.Name)
	}

	//Get pods
	fmt.Println("Pods from default namespaces")
	for _, pod := range pods.Items {
		fmt.Printf("%s\n", pod.Name)
	}

}
