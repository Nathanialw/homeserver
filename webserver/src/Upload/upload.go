package upload

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadMedia(file multipart.File, folderName string, handler *multipart.FileHeader) string {
	if handler.Filename == "" {
		fmt.Println("Image File name is empty")
		return " "
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
	fmt.Printf("systemPath: %s\n", systemPath)

	dst, err := os.Create(systemPath)
	if err != nil {
		fmt.Printf("error creating the file: %s\n", err)
		return " "
	}
	defer dst.Close()
	// Copy the uploaded file to the filesystem at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Printf("error copying the file: %s\n", err)
		return " "
	}

	return systemPath
}

func RemoveMedia(folderName string) error {
	err := os.RemoveAll(folderName)
	if err != nil {
		fmt.Printf("error removing the file: %s\n", err)
		return err
	}
	return nil
}

func UploadFolder(files map[string][]*multipart.FileHeader, baseFolder string) bool {
	// Assuming folder is a zip file, extract it
	fmt.Printf("starting to upload folder\n")
	for _, fileHeaders := range files {
		for _, fileHeader := range fileHeaders {
			fmt.Printf("Uploading fileHeader.Filename: %s\n", fileHeader.Filename)

			file, err := fileHeader.Open()
			if err != nil {
				fmt.Printf("error opening the file: %s\n", err)
				return false
			}
			defer file.Close()

			// Create the folder and subfolders
			path := filepath.Join(baseFolder, fileHeader.Filename)
			if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				fmt.Printf("error creating directories: %s\n", err)
				return false
			}

			// Create the file in the file system
			dst, err := os.Create(path)
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
		}
	}

	return true
}
