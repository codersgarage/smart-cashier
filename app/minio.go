package app

import (
	"github.com/codersgarage/smart-cashier/config"
	"github.com/minio/minio-go"
)

var spaceClient *minio.Client

func ConnectMinio() error {
	cfg := config.Minio()
	c, spaceErr := minio.New(cfg.BaseURL, cfg.Key, cfg.Secret, false)
	if spaceErr != nil {
		return spaceErr
	}

	spaceClient = c
	return nil
}

func Minio() *minio.Client {
	return spaceClient
}
