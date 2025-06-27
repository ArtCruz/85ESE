package services

import (
	"net/http"
)

type ImageUploader interface {
	UploadImage(r *http.Request) error
}

type HTTPImageUploader struct {
	APIURL string
}

func (h *HTTPImageUploader) UploadImage(r *http.Request) error {
	return UploadImage(h.APIURL, r)
}
