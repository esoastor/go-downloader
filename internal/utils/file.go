package utils

import (
	"log"
	"os"
	"path/filepath"
)

func WriteContentToFile(content []byte, filename string) {
	targetDir := filepath.Dir(filename)

	mkdirErr := os.MkdirAll(targetDir, os.ModePerm)
	if mkdirErr != nil {
		log.Fatal(mkdirErr)
	}

	var file, error = os.OpenFile(
		filename,
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 
		0644,
	)

	if error != nil {
		log.Fatal(error)
	}

	file.Write(content)
	if error := file.Close(); error != nil {
		log.Fatal(error)
	}
}

