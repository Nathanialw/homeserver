package authenticate

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"

	db "webserver/src/DB"
	serverdb "webserver/src/DB"
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
		fmt.Println("year is empty")
		return false
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
		fmt.Println("year is empty")
		return false
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

func ValidImage(folderName string, handler *multipart.FileHeader) bool {
	// Check the file type
	fileType := handler.Header.Get("Content-Type")
	fmt.Printf("/mnt/media/movies/%s/%s\n", folderName, handler.Filename)

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
	if _, err := os.Stat("/mnt/media/movies/" + folderName + "/" + handler.Filename); err == nil {
		//maybe append a number to the filename?
		fmt.Printf("file already exists: %s\n", handler.Filename)
		return false
	}
	return true
}

func ValidVideo(folderName string, handler *multipart.FileHeader) bool {
	// Check the file type
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
	if _, err := os.Stat("/mnt/media/movies/" + folderName + "/" + handler.Filename); err == nil {
		//maybe append a number to the filename?
		fmt.Printf("file already exists: %s\n", handler.Filename)
		return false
	}

	return true
}

func FormVideo(media string, r *http.Request) (multipart.File, string, *multipart.FileHeader) {
	var filename string
	var file multipart.File
	var handler *multipart.FileHeader
	var err error

	//Retrieve the file from form data
	err = r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		return file, "", handler
	}

	file, handler, err = r.FormFile(media) // "image" is the name of the file input field
	if err != nil {
		fmt.Printf("error retrieving the file, setting empty: %s\n", err)
		filename = ""
	} else {
		filename = handler.Filename
	}
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		return file, "", handler
	}

	return file, filename, handler
}

func UploadImage(file multipart.File, filename string, handler *multipart.FileHeader) bool {
	//verify image
	var imagePath string
	// var imagePath100 string
	// var imagePath400 string
	if !TextNotEmpty(filename) {
		fmt.Print("image field empty\n")
		if db.Database == serverdb.Database {
			fmt.Print("failed to add movie to content\n")
			return false
		} else {

			//need to get the UID of the book somehow
			rows, err := db.Database.Query("select title from imagePath, where uid = ?")
			if err != nil {
				fmt.Printf("error image does not exist in contentDB exists: %s\n", err)
				return false
			}
			for rows.Next() {
				rows.Scan(&imagePath)
				fmt.Print("add movie image to submit book\n")
			}
		}
	} else {
		// if either fb handler is not empty
		// if !VerifyImage(handler) {
		// 	return false
		// }
		// Create the file in the file system
		systemPath := "/mnt/media/movies/" + handler.Filename
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

		OSFile, _ := os.Open(systemPath)
		defer file.Close()

		//resize the image
		// img, _, _ := image.Decode(OSFile)
		if err != nil {
			fmt.Printf("error decoding the image: %s, %s\n", err, systemPath)
			return false
		}

		extension := filepath.Ext(systemPath)
		switch extension {
		case ".jpg", ".jpeg":
			fmt.Printf("jpeg\n")
			OSFile.Seek(0, 0) // Reset the reader to the start of the file
			img, _ := jpeg.Decode(OSFile)
			//resize the image to 400x400
			m := resize.Resize(400, 0, img, resize.Lanczos3)
			systemPath = "/mnt/media/movies/" + handler.Filename
			out, _ := os.Create(systemPath)
			defer out.Close()
			// Write the new image to the new file
			jpeg.Encode(out, m, nil)
		case ".png":
			fmt.Printf("png\n")
			OSFile.Seek(0, 0) // Reset the reader to the start of the file
			img, _ := png.Decode(OSFile)
			//resize the image to 400x400
			m := resize.Resize(400, 0, img, resize.Lanczos3)
			systemPath = "/mnt/media/movies/" + handler.Filename
			out, _ := os.Create(systemPath)
			defer out.Close()
			//write the new image to the new file
			png.Encode(out, m)
		case ".gif":
			fmt.Printf("unsupported image format: %s\n", extension)
			return false
			//OSFile.Seek(0, 0) // Reset the reader to the start of the file
			//img, err = gif.Decode(OSFile)
			////resize the image to 400x400
			//m := resize.Resize(400, 0, img, resize.Lanczos3)
			//systemPath = "../../public/assets/images/book_covers/400_" + handler.Filename
			//out, _ := os.Create(systemPath)
			//defer out.Close()
			//// Write the new image to the new file
			//gif.Encode(out, m, nil)
			//
			//m = resize.Resize(100, 0, img, resize.Lanczos3)
			//systemPath = "../../public/assets/images/book_covers/100_" + handler.Filename
			//out, _ = os.Create(systemPath)
			//gif.Encode(out, m, nil)
		default:
			fmt.Printf("unsupported image format: %s\n", extension)
			return false
		}

		//add book to database
		imagePath = "/assets/images/book_covers/" + handler.Filename
	}

	return true
}
