package main

import (
	"KubeInsight/iam/server/router"
	"KubeInsight/pkg/store"
	"log"
)

func init() {
	if err := store.InitStoreClient(); err != nil {
		log.Fatalf("Init Store Client Failed: %v", err)
	}
}

func main() {
	r := router.Router()
	r.Run()
}
