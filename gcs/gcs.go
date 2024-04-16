package gcs

import (
	"context"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var StorageClient *storage.Client
var StorageContext context.Context

func ConnectStorage() {
	// Connect to GCS
	log.Println("Connecting to GCS")
	StorageContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var err error
	if os.Getenv("ENV") == "DEV" {
		StorageClient, err = storage.NewClient(StorageContext, option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_DEV")))
	} else {
		StorageClient, err = storage.NewClient(StorageContext)
	}
	if err != nil {
		log.Println("Error connecting to GCS")
		log.Println(err)
	}
}
