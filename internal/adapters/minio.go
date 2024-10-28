package adapters

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
)

type MinioAdapter struct {
	*minio.Client
	cfg *bootstrap.Config
}

func NewMinioAdapter(cfg *bootstrap.Config) *MinioAdapter {
	return &MinioAdapter{
		cfg: cfg,
	}
}

func (a *MinioAdapter) Init() error {
	newClient, err := minio.New(a.cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(a.cfg.Minio.User, a.cfg.Minio.Password, ""),
		Secure: false,
	})

	if err != nil {
		return err
	}

	a.Client = newClient

	return nil
}
