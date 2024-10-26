package utils

import (
	"mime/multipart"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/resterr"
)

func DecodeImage(r *http.Request, cfg *bootstrap.Config) (multipart.File, *multipart.FileHeader, resterr.RestErr) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, nil, resterr.NewBadRequestError("multipart form is not valid")
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return nil, nil, resterr.NewBadRequestError("formfile is not valid")
	}
	defer file.Close()

	validateErr := ValidateImage(file, cfg)
	if validateErr != nil {
		return nil, nil, validateErr
	}

	return file, fileHeader, nil
}

func ValidateImage(f multipart.File, cfg *bootstrap.Config) resterr.RestErr {
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return resterr.NewBadRequestError("formfile is not valid")
	}

	_, _ = f.Seek(0, 0)

	mimeType := http.DetectContentType(buffer)

	for _, validType := range cfg.Image.ValidTypes {
		if mimeType == validType {
			return nil
		}
	}

	return resterr.NewBadRequestError("wrong format of image")
}
