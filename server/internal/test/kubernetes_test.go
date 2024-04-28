package test

import (
	"KubeInsight/pkg/kubernetes/client"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"testing"
)

func TestKubeClient(t *testing.T) {
	config, err := generateK8sConfig()
	if err != nil {
		t.Fatal(err)
	}
	kubeClient, err := client.NewKubernetesClient(config)
	if err != nil {
		t.Fatal(err)
	}
	pods, err := kubeClient.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}

	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}
}

func generateK8sConfig() (*rest.Config, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	// if you want to change the loading rules (which files in which order), you can do so here
	configOverrides := &clientcmd.ConfigOverrides{}
	// if you want to change override values or bind them to flags, there are methods to help you
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	return kubeConfig.ClientConfig()
}
