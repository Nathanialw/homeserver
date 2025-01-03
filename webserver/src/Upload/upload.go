package upload

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func UploadMedia(file multipart.File, folderName string, handler *multipart.FileHeader) bool {
	if handler.Filename == "" {
		fmt.Println("Image File name is empty")
		return false
	}

	// Create the folder and subfolders
	path := folderName
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	// Create the file in the file system
	systemPath := path + "/" + handler.Filename
	dst, err := os.Create(systemPath)
	if err != nil {
		fmt.Printf("error creating the file: %s\n", err)
		return false
	}
	defer dst.Close()
	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Printf("error copying the file: %s\n", err)
		return false
	}

	return true
}

func RemoveMedia(folderName string) error {
	err := os.RemoveAll(folderName)
	if err != nil {
		fmt.Printf("error removing the file: %s\n", err)
		return err
	}
	return nil
}
