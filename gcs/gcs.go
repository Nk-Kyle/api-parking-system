package gcs

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/storage"
)

var StorageClient *storage.Client

func ConnectStorage() {
	// Connect to GCS
	log.Println("Connecting to GCS")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	var err error
	StorageClient, err = storage.NewClient(ctx)
	if err != nil {
		log.Println("Error connecting to GCS")
		log.Println(err)
	}
}
