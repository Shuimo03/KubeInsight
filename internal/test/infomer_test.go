package test

import (
	"KubeInsight/pkg/kubernetes/client"
	"KubeInsight/pkg/kubernetes/informer"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestInformer(t *testing.T) {
	config, err := generateK8sConfig()
	if err != nil {
		t.Fatal(err)
	}
	kubeClient, err := client.NewKubernetesClient(config)
	if err != nil {
		t.Fatal(err)
	}

	inf := informers.NewSharedInformerFactory(kubeClient, 0)
	dl := informer.NewKubeResource(inf)

	stopCh := make(chan struct{})
	defer close(stopCh)

	go inf.Start(stopCh)
	deployments, err := dl.Deployment.ListResource()
	if !cache.WaitForCacheSync(stopCh, deployments.HasSynced) {
		t.Fatal("timed out waiting for caches to sync")
	}
	d := deployments.GetStore().List()

	printDeployments(d)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
}

func printDeployments(deployments []interface{}) {

	for _, obj := range deployments {
		d, ok := obj.(*v1.Deployment)
		if !ok {
			log.Fatalf("Error casting to Deployment object: %v", obj)
		}
		fmt.Printf("Found NameSpace: %s Deployment: %s\n", d.GetNamespace(), d.GetName())
	}
}
