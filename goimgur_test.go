package goimgur_test

import (
	"bytes"
	"testing"

	"github.com/shadyabhi/goimgur"
)

func TestUploadImage(t *testing.T) {
	t.Parallel()
	resp, err := goimgur.UploadImage("test_data/image_test.jpg")
	if err != nil {
		t.Fatalf("Error executing uploadImage: %s", err)
	}
	// Print body for debugging
	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		t.Fatalf("Error executing goimgur.UploadImage: %s", err)
	}
	resp.Body.Close()
	t.Logf("Response from uploadImage: %s", resp)
}
