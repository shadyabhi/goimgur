// Package goimgur provides helper function for uploading images to imgur.com
package goimgur

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	uploadAPIURL = "https://api.imgur.com/3/image.json"
)

// ClientID sets client_id that the developer receives after registering the app with imgur.
// Set this to avoid API rate limiting
var ClientID = "62aab02c19fde1d"

func createRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	contentType := writer.FormDataContentType()
	err = writer.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Error closing writer")
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating HTTP request")
	}
	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// uploadImage takes path to file as parameter and returns error
// in case we encounter any error while uploading the image
func uploadImage(path string) (*http.Response, error) {

	client := &http.Client{}
	request, err := createRequest(uploadAPIURL, nil, "image", path)
	if err != nil {
		return nil, errors.Wrap(err, "Error adding multipart form to request struct")
	}
	request.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", ClientID))
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Error sending POST to imgur")
	}
	return resp, nil
}

// UploadImage lets you upload iamges to imgur.
// It takes path to image file as argument
func UploadImage(path string) (*http.Response, error) {
	resp, err := uploadImage(path)
	if err != nil {
		return nil, errors.Wrap(err, "Error calling UploadImage")
	}
	return resp, nil
}
