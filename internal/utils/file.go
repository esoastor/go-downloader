package utils

import (
	"os"
	"path/filepath"
)

func WriteContentToFile(content []byte, filename string) error {
	targetDir := filepath.Dir(filename)

	mkdirErr := os.MkdirAll(targetDir, os.ModePerm)
	if mkdirErr != nil {
		return mkdirErr
	}

	var file, error = os.OpenFile(
		filename,
		os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 
		0644,
	)

	if error != nil {
		return error
	}

	file.Write(content)
	if error := file.Close(); error != nil {
		return error
	}

	return nil
}

