package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shadyabhi/goimgur"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide file path as first argument")
	}

	path := os.Args[1]
	resp, err := goimgur.UploadImage(path)
	if err != nil {
		log.Fatal(err)
	}
	body, err := goimgur.ParseBody(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
