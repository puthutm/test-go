package config

import (
	"log"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinio(cnf *Config) *minio.Client {
	portString := strconv.Itoa(cnf.Minio.Port)
	client, err := minio.New(cnf.Minio.Endpoint+":"+portString, &minio.Options{
		Creds:  credentials.NewStaticV4(cnf.Minio.AccessKeyID, cnf.Minio.SecretAccessKey, ""),
		Secure: cnf.Minio.UseSSL,
	})
	if err != nil {
		log.Fatalln("Error creating MinIO client:", err)
	}
	return client
}
