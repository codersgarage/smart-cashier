package services

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/config"
	"github.com/minio/minio-go"
	"io"
)

func UploadToMinio(fileName, contentType string, reader io.Reader, size int64) error {
	conn := app.Minio()
	cfg := config.Minio()
	_, errP := conn.PutObject(cfg.Bucket, fileName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if errP != nil {
		return errP
	}
	return nil
}

func ServeFromMinio(fileName string) (string, error) {
	conn := app.Minio()
	cfg := config.Minio()
	url, err := conn.PresignedGetObject(cfg.Bucket, fileName, cfg.SignDuration, nil)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func ServeAsStreamFromMinio(fileName string) (*minio.Object, error) {
	conn := app.Minio()
	cfg := config.Minio()
	o, err := conn.GetObject(cfg.Bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return o, nil
}
