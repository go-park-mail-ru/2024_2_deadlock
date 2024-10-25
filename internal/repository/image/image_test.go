package image

// import (
// 	"bytes"
// 	"context"
// 	"encoding/base64"
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"time"

// 	"github.com/minio/minio-go/v7"

// 	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
// 	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
// 	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
// )

// func TestOK(t *testing.T) {
// 	t.Log("YES")
// 	absPath, _ := filepath.Abs("../../../dev.yaml")
// 	cfg, err := bootstrap.Setup(absPath)
// 	if err != nil {
// 		t.Errorf("error reading config")
// 	}
// 	t.Log("password", cfg.Minio.Password)
// 	t.Log("port", cfg.Minio.Endpoint)
// 	t.Log("user", cfg.Minio.User)

// 	minioAdapter := adapters.NewMinioAdapter(cfg)
// 	err = minioAdapter.Init()
// 	if err != nil {
// 		t.Errorf("error initializing adaptor minio")
// 	}

// 	ctx := context.Background()

// 	minioRepo := NewRepository(minioAdapter)
// 	err = minioRepo.Init(ctx, "articlephotos")
// 	if err != nil {
// 		t.Errorf("error initializing bucket %v", err)
// 	}

// 	// ------------------------------------------------------------------

// 	file, err := os.Open("sea.jpg")
// 	if err != nil {
// 		t.Errorf("error while opeining file")
// 	}
// 	defer file.Close()

// 	byteFile, err := io.ReadAll(file)
// 	if err != nil {
// 		t.Errorf("unable to turn file into bytes")
// 	}

// 	// ------------------------------------------------------------------

// 	encoded := base64.StdEncoding.EncodeToString(byteFile)
// 	imageToUpload := &domain.ImageData{Image: encoded}
// 	imageUploadData, err := minioRepo.PutImage(ctx, imageToUpload)
// 	if err != nil {
// 		t.Errorf("cant upload image %v", err)
// 	}
// 	t.Log("created image", imageUploadData)
// 	url, err := minioRepo.GetImage(ctx, imageUploadData.ID)
// 	if err != nil {
// 		t.Errorf("cant get image %v", err)
// 	}
// 	t.Log("get image ", url)
// }

// func TestDeleteOK(t *testing.T) {
// 	absPath, _ := filepath.Abs("../../../../dev.yaml")
// 	cfg, err := bootstrap.Setup(absPath)
// 	if err != nil {
// 		t.Errorf("error reading config")
// 	}
// 	t.Log("password", cfg.Minio.Password)
// 	t.Log("port", cfg.Minio.Endpoint)
// 	t.Log("user", cfg.Minio.User)

// 	minioAdapter := adapters.NewMinioAdapter(cfg)
// 	err = minioAdapter.Init()
// 	if err != nil {
// 		t.Errorf("error initializing adaptor minio")
// 	}
// 	ctx := context.Background()
// 	minioRepo := NewRepository(minioAdapter)
// 	err = minioRepo.Init(ctx, "articlephotos")
// 	if err != nil {
// 		t.Errorf("error initializing bucket %v", err)
// 	}
// 	ImageID := "98a44f80-483f-4b42-af58-831a30282936"
// 	err = minioRepo.DeleteImage(ctx, domain.ImageID(ImageID))
// 	if err != nil {
// 		t.Errorf("unable to delete image %v", err)
// 	}
// 	t.Log("image deleted !!!")
// }

// func TestGetOK(t *testing.T) {
// 	absPath, _ := filepath.Abs("../../../../dev.yaml")
// 	cfg, err := bootstrap.Setup(absPath)
// 	if err != nil {
// 		t.Errorf("error reading config")
// 	}
// 	t.Log("password", cfg.Minio.Password)
// 	t.Log("port", cfg.Minio.Endpoint)
// 	t.Log("user", cfg.Minio.User)

// 	minioAdapter := adapters.NewMinioAdapter(cfg)
// 	err = minioAdapter.Init()
// 	if err != nil {
// 		t.Errorf("error initializing adaptor minio")
// 	}
// 	ctx := context.Background()
// 	minioRepo := NewRepository(minioAdapter)
// 	err = minioRepo.Init(ctx, "articlephotos")
// 	if err != nil {
// 		t.Errorf("error initializing bucket %v", err)
// 	}
// 	ImageID := "98a44f80-483f-4b42-af58-831a30282936"
// 	ReqID := domain.ImageID(ImageID)
// 	t.Log("reqID", ReqID)
// 	val, err := minioRepo.GetImage(ctx, ReqID)
// 	if err != nil {
// 		t.Errorf("unable to delete image %v", err)
// 	}
// 	t.Log("image got !!!", val)
// }

// func TestUpdateOK(t *testing.T) {
// 	t.Log("YES")
// 	absPath, _ := filepath.Abs("../../../../dev.yaml")
// 	cfg, err := bootstrap.Setup(absPath)
// 	if err != nil {
// 		t.Errorf("error reading config")
// 	}
// 	t.Log("password", cfg.Minio.Password)
// 	t.Log("port", cfg.Minio.Endpoint)
// 	t.Log("user", cfg.Minio.User)

// 	minioAdapter := adapters.NewMinioAdapter(cfg)
// 	err = minioAdapter.Init()
// 	if err != nil {
// 		t.Errorf("error initializing adaptor minio")
// 	}

// 	ctx := context.Background()

// 	minioRepo := NewRepository(minioAdapter)
// 	err = minioRepo.Init(ctx, "articlephotos")
// 	if err != nil {
// 		t.Errorf("error initializing bucket %v", err)
// 	}

// 	// ------------------------------------------------------------------

// 	file, err := os.Open("sea.jpg")
// 	if err != nil {
// 		t.Errorf("error while opeining file")
// 	}
// 	defer file.Close()

// 	byteFile, err := io.ReadAll(file)
// 	if err != nil {
// 		t.Errorf("unable to turn file into bytes")
// 	}

// 	encoded := base64.StdEncoding.EncodeToString(byteFile)
// 	imageToUpload := &domain.ImageData{Image: encoded}
// 	image, err := base64.StdEncoding.DecodeString(imageToUpload.Image)
// 	if err != nil {
// 		t.Errorf("unable to decode string from base64 format %v", err)
// 	}

// 	xfile := bytes.NewReader(image)
// 	imageID := "b80f7ace-4b28-4122-8c31-6e06397a887a"

// 	_, err = minioRepo.MinioAdapter.PutObject(ctx, minioRepo.BucketName, imageID, xfile, int64(len(image)),
// 		minio.PutObjectOptions{})
// 	if err != nil {
// 		t.Errorf("unable to upload photo % v", err)
// 	}

// 	url, err := minioRepo.MinioAdapter.PresignedGetObject(ctx, minioRepo.BucketName, imageID, time.Second*60*60*24, nil)
// 	if err != nil {
// 		t.Errorf("unable to get url of uploaded photo %v", err)
// 	}

// 	t.Log("newfile", imageID, url)
// }
