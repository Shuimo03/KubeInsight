package test

import (
	"KubeInsight/kubernetes/client"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"testing"
)

func TestKubeClient(t *testing.T) {
	kubeClient, err := client.NewKubernetesClient()
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
