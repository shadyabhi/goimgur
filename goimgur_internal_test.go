package goimgur

import "testing"

func TestUploadImage(t *testing.T) {
	err := uploadImage("test_data/image_test.jpg")
	if err != nil {
		t.Fatalf("Error executing uploadImage: %s", err)
	}
}
