package main

import (
	"os"

	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeAPIController interface {
	ServerVersion() (version.Info, error)
}

type kubeAPIControllerImpl struct {
	clientset *kubernetes.Clientset
}

func InitKubeAPIController() (KubeAPIController, error) {
	// First check for the existence of a KUBECONFIG environment variable, and
	// if it exists, use that to create a client config. If it doesn't exist,
	// assume we're running in-cluster and try to create an in-cluster config.

	var config *rest.Config
	var err error

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig != "" {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &kubeAPIControllerImpl{clientset: clientset}, nil
}

func (c *kubeAPIControllerImpl) ServerVersion() (version.Info, error) {
	info, err := c.clientset.Discovery().ServerVersion()
	if err != nil {
		return version.Info{}, err
	}
	return *info, nil
}
