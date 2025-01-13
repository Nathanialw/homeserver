package lantv

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

// func authenticateSeason(w http.ResponseWriter, r *http.Request) (bool, Series) {
// 	//drop in a folder
// 	//upload the folder
// 	//save the files in th subfolders, maintain structure
// 	//save it into the struct

// 	var series Series

// 	var folderName string = "/mnt/media/tv/" + currentSeriesTitle + "_" + currentSeriesSubtitle

// 	fmt.Printf("folderName: %s\n", folderName)

// 	// movie.Image = folderName + "/" + imageHandler.Filename
// 	// folder, videoHandler := authenticate.FormMediaFolder(r)

// 	fmt.Printf("parse\n")

// 	err := r.ParseMultipartForm(10 << 20) // 10 MB
// 	if err != nil {
// 		fmt.Printf("error parsing the form: %s\n", err)
// 		return false, series
// 	}

// 	fmt.Printf("uploading\n")
// 	if !upload.UploadFolder(r.MultipartForm.File, folderName) {
// 		return false, series
// 	}

// 	for _, fileHeaders := range r.MultipartForm.File {
// 		season := Season{}
// 		season.SeasonNum, _ = strconv.Atoi(r.FormValue("season"))
// 		series.Seasons = append(series.Seasons, season)

// 		for i := 0; i < len(fileHeaders); i++ {
// 			var episode Episode
// 			episode.Title = fileHeaders[i].Filename
// 			episode.Path = folderName + "/" + episode.Title
// 			series.Seasons[0].Episodes = append(series.Seasons[0].Episodes, episode)
// 			fmt.Printf("episode name: %s, number: %d\n", fileHeaders[i].Filename, i)
// 		}
// 	}

// 	// authenticate.ValidVideo(folderName, videoHandler)

// 	// imageFile.Close()
// 	return true, series
// }
