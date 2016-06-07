package goimgur

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// Imgur API details
const (
	uploadAPIUrl = "https://api.imgur.com/3/image.json"
	clientID     = "62aab02c19fde1d"
)

func getFile(path string) (string, error) {
	imgFile, err := ioutil.ReadFile(path)
	if err != nil {
		return "", errors.Wrap(err, "Error opening image file")
	}

	base64 := base64.StdEncoding.EncodeToString(imgFile)
	return base64, nil
}

// uploadImage takes path to file as parameter and returns error
// in case we encounter any error while uploading the image
func uploadImage(path string) error {
	base64, err := getFile(path)
	if err != nil {
		return errors.Wrap(err, "Error converting image file to base64")
	}

	client := &http.Client{}
	form := url.Values{}
	form.Set("image", base64)
	req, err := http.NewRequest("GET", uploadAPIUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return errors.Wrap(err, "Error generating request to Imgur API")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", clientID))

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Error doing POST to imgur")
	}
	fmt.Printf("\n\n\n %#v", resp)
	return nil
}
