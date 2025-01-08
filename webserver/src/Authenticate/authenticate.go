package authenticate

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func TextNotEmpty(title string) bool {
	if title == "" {
		fmt.Printf("%s is empty\n", title)
		return false
	}
	return true
}

func ValidText(text string) bool {
	return TextNotEmpty(text)
}

// func ValidTitle(title string) bool {
// 	return TextNotEmpty(title)
// }

// func ValidDirector(title string) bool {
// 	return TextNotEmpty(title)
// }

// func Valid(title string) bool {
// 	return TextNotEmpty(title)
// }

func ValidLength(length string) bool {
	if length == "" {
		fmt.Println("length is empty") //length not set
		return true
	}
	for _, c := range length {
		if c < '0' || c > '9' {
			fmt.Println("year is not a number")
			return false
		}
	}
	if length == "0" {
		fmt.Println("length of 0 is irrational")
		return false
	}
	num, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		fmt.Printf("error converting string to int in authenticate.ValidLength(): %s\n", err)
		return false
	}
	if num > 1000 {
		fmt.Printf("The length is longer than the max of 1000 minutes: %s\n", length)
		return false
	}
	return true
}

func ValidYear(year string) bool {
	if year == "" {
		fmt.Println("year is empty") //year not set
		return true
	}
	for _, c := range year {
		if c < '0' || c > '9' {
			fmt.Println("year is not a number")
			return false
		}
	}
	if len(year) > 4 {
		fmt.Println("year is more than 4 characters")
		return false
	}
	if year == "0" {
		fmt.Println("there is no year 0 BC or 0 AD")
		return false
	}
	if year > "2024" {
		fmt.Println("the max year is 2024 AD")
		return false
	}
	return true
}

func ValidEra(era string, year string) bool {
	if era == "" {
		fmt.Println("year is empty") //era not set
		return true
	}
	if era != "BC" && era != "AD" {
		fmt.Println("era is not BC or AD")
		return false
	}
	if era == "AD" && year > "2024" {
		fmt.Println("the max year is 2024 AD")
		return false
	}
	return true
}

func ValidImage(filePath string, handler *multipart.FileHeader) bool {
	// Check the file type
	// fmt.Printf("/mnt/media/movies/%s/%s\n", folderName, handler.Filename)
	if handler.Filename == "" {
		// fmt.Printf("File name is empty: %s\n", handler.Filename)
		return true
	}

	fileType := handler.Header.Get("Content-Type")

	switch fileType {
	case "image/jpeg", "image/jpg", "image/png":
		break
	default:
		fmt.Println("File is not an image")
		return false
	}
	if handler.Size > 4*1024*1024 {
		fmt.Printf("File is too large (max 4MB) %d\n", handler.Size)
		return false
	}
	//check if the file is an image
	if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
		fmt.Printf("file is not an image: %s\n", handler.Header.Get("Content-Type"))
		return false
	}
	//check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		//maybe append a number to the filename?
		fmt.Printf("file already exists: %s\n", handler.Filename)
		return false
	}
	return true
}

func ValidVideo(folderName string, handler *multipart.FileHeader) bool {
	// // Check the file type
	// fileType := handler.Header.Get("Content-Type")

	// switch fileType {
	// case "image/jpeg", "image/jpg", "image/png":
	// 	break
	// default:
	// 	fmt.Println("File is not an image")
	// 	return false
	// }
	// if handler.Size > 4*1024*1024 {
	// 	fmt.Printf("File is too large (max 4MB) %d\n", handler.Size)
	// 	return false
	// }
	// //check if the file is an image
	// if !strings.Contains(handler.Header.Get("Content-Type"), "image") {
	// 	fmt.Printf("file is not an image: %s\n", handler.Header.Get("Content-Type"))
	// 	return false
	// }
	// //check if the file already exists
	// if _, err := os.Stat("/mnt/media/movies/" + folderName + "/" + handler.Filename); err == nil {
	// 	//maybe append a number to the filename?
	// 	fmt.Printf("file already exists: %s\n", handler.Filename)
	// 	return false
	// }

	return true
}

func FormMedia(media string, r *http.Request) (multipart.File, *multipart.FileHeader) {
	var file multipart.File
	var handler *multipart.FileHeader
	var err error

	//Retrieve the file from form data
	err = r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		return file, handler
	}

	file, handler, err = r.FormFile(media) // "image" is the name of the file input field
	if err != nil {
		fmt.Printf("error retrieving the file, setting empty: %s\n", err)
	}
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		return file, handler
	}

	return file, handler
}

func FormMediaFolder(r *http.Request) (map[string][]*multipart.FileHeader, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		return nil, err
	}

	if r.MultipartForm == nil || len(r.MultipartForm.File) == 0 {
		fmt.Println("Form is empty")
		return nil, fmt.Errorf("form is empty")
	}

	return r.MultipartForm.File, nil
}
