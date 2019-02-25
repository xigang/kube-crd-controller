package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	crdclientset "github.com/xigang/kube-crd-controller/pkg/client/clientset/versioned"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master     = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	if *master == "" || *kubeconfig == "" {
		return
	}

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kubeconfig)
	if err != nil {
		logrus.Fatalf("Error building kubeconfig: %+v", err)
	}

	customResourceClient, err := crdclientset.NewForConfig(cfg)
	if err != nil {
		logrus.Fatalf("Error building custome resource clientset: %+v", err)
	}

	list, err := customResourceClient.ExampleV1().MyCustomResources("default").List(metav1.ListOptions{})
	if err != nil {
		logrus.Fatalf("Error listing all custom resource: %+v", err)
	}

	for _, resource := range list.Items {
		fmt.Printf("custom resource name: %+v", resource)
	}
}
