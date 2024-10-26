package image

import (
	"bytes"
	"context"
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
)

type Repository struct {
	pg.MinioRepo
	BucketName string
}

func NewRepository(adapter *adapters.MinioAdapter) *Repository {
	return &Repository{
		MinioRepo: pg.MinioRepo{
			MinioAdapter: adapter,
		},
	}
}

func (r *Repository) Init(ctx context.Context, bucketName string) error {
	bucketExists, err := r.MinioAdapter.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if bucketExists {
		r.BucketName = bucketName
		return nil
	}

	err = r.MinioAdapter.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}

	r.BucketName = bucketName

	return nil
}

func (r *Repository) PutImage(ctx context.Context, data *domain.ImageData) (*domain.ImageUploadInfo, error) {
	image, err := base64.StdEncoding.DecodeString(data.Image)
	if err != nil {
		return nil, interr.NewInternalError(err, "unable to decode string from base64 format")
	}

	file := bytes.NewReader(image)
	imageID := uuid.New().String()

	_, err = r.MinioAdapter.PutObject(ctx, r.BucketName, imageID, file, int64(len(image)),
		minio.PutObjectOptions{})
	if err != nil {
		return nil, interr.NewInternalError(err, "unable to upload photo")
	}

	url, err := r.MinioAdapter.PresignedGetObject(ctx, r.BucketName, imageID, time.Second*60*60*24, nil)
	if err != nil {
		return nil, interr.NewInternalError(err, "unable to get url of uploaded photo")
	}

	imageUploadInfo := new(domain.ImageUploadInfo)
	imageUploadInfo.ID = domain.ImageID(imageID)
	imageUploadInfo.URL = domain.ImageURL(url.String())

	return imageUploadInfo, nil
}

func (r *Repository) GetImage(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error) {
	_, err := r.MinioAdapter.StatObject(ctx, r.BucketName, string(imageID), minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return "", interr.NewNotFoundError("current imageID does not exist")
		} else {
			return "", interr.NewInternalError(err, "error while viewing image by imageID")
		}
	}

	url, err := r.MinioAdapter.PresignedGetObject(ctx, r.BucketName, string(imageID), time.Second*60*60*24, nil)
	if err != nil {
		return "", interr.NewInternalError(err, "unable to get url of uploaded photo")
	}

	imageURL := domain.ImageURL(url.String())

	return imageURL, nil
}

func (r *Repository) DeleteImage(ctx context.Context, imageID domain.ImageID) error {
	err := r.MinioAdapter.RemoveObject(ctx, r.BucketName, string(imageID), minio.RemoveObjectOptions{})
	if err != nil {
		return interr.NewInternalError(err, "unable to delete image")
	}

	return nil
}
