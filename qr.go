package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <file-name>", os.Args[0])
	}

	org_file := os.Args[1]
	content, err := ioutil.ReadFile(org_file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	base64_content := base64.StdEncoding.EncodeToString(content)

	data_uri := fmt.Sprintf("data:text/html;base64,%s", base64_content)

	qr_file_name := "qr.png" // i need to change this

	err = qrcode.WriteFile(data_uri, qrcode.Medium, 256, qr_file_name)

	if err != nil {
		log.Fatalf("failed to generate qr code: %v", err)
	}

	fmt.Printf("qr code generated and saved as %s\n", qr_file_name)
}
