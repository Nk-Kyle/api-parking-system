package gcs

import (
	"api-parking-system/utils"
	"context"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
)

type FileUploader struct {
	Client     *storage.Client
	ProjectID  string
	BucketName string
}

func (c *FileUploader) UploadFile(file multipart.File, object string) (string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// use current timestamp as object name + random string
	object = object + strconv.FormatInt(time.Now().Unix(), 10) + utils.GenerateRandomString(5)
	wc := c.Client.Bucket(c.BucketName).Object(object).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return wc.Attrs().MediaLink, nil
}
