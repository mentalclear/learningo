package main

import (
	"archive/zip"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	fmt.Println("This is err:", err)
	fmt.Println("This is zip:", zip.ErrFormat)
	if err == zip.ErrFormat {
		fmt.Println("Told you so")
	}
}
