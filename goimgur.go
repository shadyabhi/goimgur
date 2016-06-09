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

// Imgur API details
const (
	uploadAPIUrl = "https://api.imgur.com/3/image.json"
	clientID     = "62aab02c19fde1d"
)

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
func uploadImage(path string) (*bytes.Buffer, error) {

	client := &http.Client{}
	request, err := createRequest(uploadAPIUrl, nil, "image", "test_data/image_test.jpg")
	if err != nil {
		return nil, errors.Wrap(err, "Error adding multipart form to request struct")
	}
	request.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", clientID))
	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "Error sending POST to imgur")
	}
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading from response body")
	}
	resp.Body.Close()
	return body, nil
}
