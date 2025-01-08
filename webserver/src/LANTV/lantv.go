package lantv

import (
	"fmt"
	"net/http"
	"strconv"
	authenticate "webserver/src/Authenticate"
	content "webserver/src/Content"
	upload "webserver/src/Upload"
	user "webserver/src/User"

	"github.com/julienschmidt/httprouter"
)

var currentSeriesTitle string
var currentSeriesSubtitle string

type Episode struct {
	Uid        int
	Series     string
	EpisodeNum int
	Season     int
	Title      string
	Subtitle   string
	Image      string
	Synopsis   string
	Path       string
	Back       string
}

type Season struct {
	Uid       int
	Title     string
	Image     string
	SeasonNum int
	Synopsis  string
	Path      string
	Episodes  []Episode
	Back      string
}

type Series struct {
	Uid      int
	Title    string
	Subtitle string
	Writer   string
	Image    string
	Synopsis string
	Path     string
	Seasons  []Season
	Back     string
}

type List struct {
	User     user.Session
	NotEmpty bool
	Series   []Series
}

//list the series
func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createSeriesDB()
	createEpisodesDB()

	var err error
	var data List

	data.Series, err = getAll()
	if len(data.Series) > 0 {
		data.NotEmpty = true
		fmt.Printf("series found %s.\n", data.Series[0].Title)
	} else {
		data.NotEmpty = false
		fmt.Printf("none found\n")
	}

	content.GenerateHTML(w, data, "LANTV", "home")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//list the seasons of the selected series
func ShowSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createSeriesDB()
	createEpisodesDB()

	var err error
	var data Series
	var episodes []Episode

	//need to mkae sure th "movieID" actually exists so it can 404 if it doesn't
	data, _ = RetrieveSeriesFromDB(p.ByName("seriesID"))
	fmt.Printf("title is %s\n", data.Title)

	currentSeriesTitle = data.Title
	currentSeriesSubtitle = data.Subtitle

	episodes, _ = RetrieveEpisodesFromDB(data.Title)
	fmt.Printf("number of episodes: %d\n", len(episodes))

	data = OrganizeIntoSeasons(data, episodes)
	fmt.Printf("number of seasons: %d\n", len(data.Seasons))

	data.Back = "/tv/"
	// /tv/:seriesID/:seasonNum
	content.GenerateHTML(w, data, "LANTV", "series")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//list the episodes of the selected season
func ShowSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createSeriesDB()
	createEpisodesDB()

	var err error
	var data Series
	var episodes []Episode

	fmt.Printf("Series: %s\n", p.ByName("seriesID"))
	fmt.Printf("Season %s\n", p.ByName("seasonNum"))

	//need to mkae sure th "movieID" actually exists so it can 404 if it doesn't
	data, err = RetrieveSeriesFromDB(p.ByName("seriesID"))

	//organize the episodes by season
	episodes, _ = RetrieveEpisodesFromDB(data.Title)
	fmt.Printf("number of episodes: %d\n", len(episodes))

	data = OrganizeIntoSeasons(data, episodes)
	fmt.Printf("number of seasons: %d\n", len(data.Seasons))

	num, _ := strconv.Atoi(p.ByName("seasonNum"))
	data.Seasons[num-1].Back = "/tv/" + p.ByName("seriesID")
	content.GenerateHTML(w, data.Seasons[num-1], "LANTV", "seasons")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ShowEpisode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createSeriesDB()
	createEpisodesDB()

	var err error
	var data Series
	var episodes []Episode

	fmt.Printf("Series: %s\n", p.ByName("seriesID"))
	fmt.Printf("Season %s\n", p.ByName("seasonNum"))
	fmt.Printf("Episode: %s\n", p.ByName("episodeNum"))

	//need to mkae sure th "movieID" actually exists so it can 404 if it doesn't
	data, err = RetrieveSeriesFromDB(p.ByName("seriesID"))

	//organize the episodes by season
	episodes, _ = RetrieveEpisodesFromDB(data.Title)
	fmt.Printf("number of episodes: %d\n", len(episodes))

	data = OrganizeIntoSeasons(data, episodes)
	fmt.Printf("number of seasons: %d\n", len(data.Seasons))

	seasonNum, _ := strconv.Atoi(p.ByName("seasonNum"))
	episodeNum, _ := strconv.Atoi(p.ByName("episodeNum"))

	data.Seasons[seasonNum-1].Episodes[episodeNum-1].Back = "/tv/" + p.ByName("seriesID") + "/" + p.ByName("seasonNum")
	content.GenerateHTML(w, data.Seasons[seasonNum-1].Episodes[episodeNum-1], "LANTV", "episode")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func authenticateSeries(w http.ResponseWriter, r *http.Request) (bool, Series) {
	//drop in a folder
	//upload the folder
	//save the files in th subfolders, maintain structure
	//save it into the struct

	var series Series
	var success bool = false

	series.Title = r.FormValue("title")
	series.Subtitle = r.FormValue("subtitle")
	currentSeriesTitle = series.Title
	currentSeriesSubtitle = series.Subtitle

	// movie.Image = folderName + "/" + imageHandler.Filename
	// folder, videoHandler := authenticate.FormMediaFolder(r)

	if authenticate.ValidText(series.Title) {
		success = true
	} else {
		return false, series
	}
	// authenticate.ValidVideo(folderName, videoHandler)

	// imageFile.Close()
	return success, series
}

func authenticateSeason(w http.ResponseWriter, r *http.Request) (bool, Series) {
	//drop in a folder
	//upload the folder
	//save the files in th subfolders, maintain structure
	//save it into the struct

	var series Series

	var folderName string = "/mnt/media/tv/" + currentSeriesTitle + "_" + currentSeriesSubtitle

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
		season.SeasonNum, _ = strconv.Atoi(r.FormValue("season"))
		series.Seasons = append(series.Seasons, season)

		for i := 0; i < len(fileHeaders); i++ {
			var episode Episode
			episode.Title = fileHeaders[i].Filename
			episode.Path = folderName + "/" + episode.Title
			series.Seasons[0].Episodes = append(series.Seasons[0].Episodes, episode)
			fmt.Printf("episode name: %s, number: %d\n", fileHeaders[i].Filename, i)
		}
	}

	// authenticate.ValidVideo(folderName, videoHandler)

	// imageFile.Close()
	return true, series
}

func AddSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data List
	// user.Session.LoggedIn = LoginStatus(r)
	// user.Session.Admin = AdminStatus(r)
	content.GenerateHTML(w, data, "LANTV", "addseries")
}

func SubmitSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	success, series := authenticateSeries(w, r)
	if success {
		insertSeriesIntoDB(series)
		fmt.Print("added\n")
	} else {
		fmt.Print("not added\n")
	}

	http.Redirect(w, r, "/tv", http.StatusSeeOther)
}

func AddSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data List
	// user.Session.LoggedIn = LoginStatus(r)
	// user.Session.Admin = AdminStatus(r)
	content.GenerateHTML(w, data, "LANTV", "addseason")
}

func SubmitSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	fmt.Printf("series title: %s\n", p.ByName("seriesID"))

	success, series := authenticateSeason(w, r)
	if success {
		insertSeasonIntoDB(series)
		fmt.Print("added\n")
	} else {
		fmt.Print("not added\n")
	}

	http.Redirect(w, r, "/tv/"+p.ByName("seriesID"), http.StatusSeeOther)
}

func getAll() (series []Series, err error) {

	series, _ = getAllSeries()

	fmt.Println("retrieving all")
	return
}

func OrganizeIntoSeasons(series Series, episodes []Episode) Series {
	// get number of seasons
	numSeasons := 0
	for _, episode := range episodes {
		if episode.Season > numSeasons {
			numSeasons = episode.Season
		}
		if len(series.Seasons) < episode.Season {
			series.Seasons = append(series.Seasons, Season{})
			series.Seasons[episode.Season-1].SeasonNum = episode.Season
			series.Seasons[episode.Season-1].Title = series.Title
		}
		series.Seasons[episode.Season-1].Episodes = append(series.Seasons[episode.Season-1].Episodes, episode)
	}

	//organize the episodes in each season
	for _, season := range series.Seasons {
		var organizedSeason = season

		for _, episode := range season.Episodes {
			organizedSeason.Episodes[episode.EpisodeNum-1] = episode
		}

		season = organizedSeason
	}

	//get the episodes
	fmt.Printf("num seasons: %d\n", len(series.Seasons))

	for _, season := range series.Seasons {
		fmt.Printf("season title: %s\n", series.Title)
		fmt.Printf("season number: %d\n", season.SeasonNum)
		for _, episode := range season.Episodes {
			fmt.Printf("name: %s\n", episode.Title)
		}
	}

	//organize the episodes by season

	return series
}
