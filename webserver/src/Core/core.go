package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	authenticate "webserver/src/Authenticate"
	content "webserver/src/Content"
	upload "webserver/src/Upload"
	user "webserver/src/User"

	"github.com/julienschmidt/httprouter"
)

type Series struct {
	ID    string
	Title string
	Image string
}

type List struct {
	User     user.Session
	NotEmpty bool
	Media    []Series

	// routing
	Back         string
	Add          string
	Submit       string
	Route        string
	PreviewRoute string
}

//list the series
func Home(w http.ResponseWriter, add string, module string, moduleType string) {

	var err error
	var data List

	data.Media, err = getAllSeries(moduleType)

	if len(data.Media) > 0 {
		data.NotEmpty = true
		fmt.Printf("found %s.\n", data.Media[0].Title)
	} else {
		data.NotEmpty = false
		fmt.Printf("none found\n")
	}

	data.Back = "/"
	data.Add = add

	content.GenerateHTML(w, data, module, "home")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Add(w http.ResponseWriter, back string, module string, template string, submit string, route string, previewRoute string) {

	var data List
	// user.Session.LoggedIn = LoginStatus(r)
	// user.Session.Admin = AdminStatus(r)

	data.Back = back
	data.Add = ""
	data.Submit = submit
	data.Route = route
	data.PreviewRoute = previewRoute

	content.GenerateHTML(w, data, module, template)
}

func UpdateSearch(w http.ResponseWriter, r *http.Request, dbname string) {

	query := r.FormValue("query")
	results := Search_Series(dbname, query)

	var response []byte
	var err error

	if len(results) != 0 {
		// Create response as JSON
		response, err = json.Marshal(results)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
			return
		}
	} else {
		response = []byte(`""`)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func SubmitSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func SubmitFile(w http.ResponseWriter, r *http.Request, pathSuffix string) string {
	var folderpath string = "/mnt/media" + pathSuffix

	videoFile, videoHandler := authenticate.FormMedia("media", r)
	//upload the file
	path := upload.UploadMedia(videoFile, folderpath, videoHandler)
	videoFile.Close()

	return path
}

func Install() {
	mediaTypes := [][]string{
		{"tv", "tvSeries, tvMiniSeries"},
		{"movies", "movie, tvMovie, video"},
	}

	for _, mediaType := range mediaTypes {
		cmd := exec.Command("python3", "../scripts/convert_to_db.py", mediaType[0], mediaType[1])
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Failed to run command: %v\n", err)
		}

		fmt.Printf("output: %s\n", output)
	}
}
