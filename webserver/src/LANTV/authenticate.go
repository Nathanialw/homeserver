package lantv

import (
	"fmt"
	"net/http"
	"strconv"
	upload "webserver/src/Upload"
)

// func authenticateSeries(w http.ResponseWriter, r *http.Request) (bool, Series) {
// 	//drop in a folder
// 	//upload the folder
// 	//save the files in th subfolders, maintain structure
// 	//save it into the struct

// 	var series Series
// 	var success bool = false

// 	series.Title = r.FormValue("title")
// 	series.Subtitle = r.FormValue("subtitle")
// 	currentSeriesTitle = series.Title
// 	currentSeriesSubtitle = series.Subtitle

// 	// movie.Image = folderName + "/" + imageHandler.Filename
// 	// folder, videoHandler := authenticate.FormMediaFolder(r)

// 	if authenticate.ValidText(series.Title) {
// 		success = true
// 	} else {
// 		return false, series
// 	}
// 	// authenticate.ValidVideo(folderName, videoHandler)

// 	// imageFile.Close()
// 	return success, series
// }

func authenticateSeason(w http.ResponseWriter, r *http.Request, id string, seasonNum string) (bool, Series) {
	//drop in a folder
	//upload the folder
	//save the files in th subfolders, maintain structure
	//save it into the struct

	var series Series
	series.ID = id

	var folderName string = "/mnt/media/tv/" + id + "/" + seasonNum

	fmt.Printf("folderName: %s\n", folderName)

	// movie.Image = folderName + "/" + imageHandler.Filename
	// folder, videoHandler := authenticate.FormMediaFolder(r)

	fmt.Printf("parse\n")

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		fmt.Printf("error parsing the form: %s\n", err)
		return false, series
	}

	fmt.Printf("uploading\n")
	if !upload.UploadFolder(r.MultipartForm.File, folderName) {
		return false, series
	}

	for _, fileHeaders := range r.MultipartForm.File {
		season := Season{}
		season.SeasonNum, _ = strconv.Atoi(seasonNum)

		for i := 0; i < len(fileHeaders); i++ {
			var episode Episode
			episode.Title = fileHeaders[i].Filename
			episode.Path = folderName + "/" + episode.Title
			season.Episodes = append(season.Episodes, episode)
			fmt.Printf("episode name: %s, number: %d\n", fileHeaders[i].Filename, i)
		}

		series.Seasons = append(series.Seasons, season)
		fmt.Printf("season number: %s, number of episodes: %d, number of seasons: %d\n", seasonNum, len(season.Episodes), len(series.Seasons))
	}

	// authenticate.ValidVideo(folderName, videoHandler)

	// imageFile.Close()
	return true, series
}
