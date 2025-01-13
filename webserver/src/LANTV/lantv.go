package lantv

import (
	"fmt"
	"net/http"
	content "webserver/src/Content"
	core "webserver/src/Core"

	"github.com/julienschmidt/httprouter"
)

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
	Active     bool
}

type Season struct {
	Uid       int
	Title     string
	Image     string
	SeasonNum int
	Synopsis  string
	Path      string
	Episodes  []Episode
	Active    bool
}

type Series struct {
	Uid         int
	ID          string
	Title       string
	Subtitle    string
	Writer      string
	ReleaseDate string
	Runtime     string
	Rating      string
	Genres      string
	GenresList  []string
	Ratings     string
	NumImages   string
	Review      string
	Image       string
	NumSeasons  string
	Synopsis    string
	Path        string
	Seasons     []Season
	Back        string
	Add         string
}

var SeriesData Series
var currentSeriesTitle string

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createSeriesDB()
	createEpisodesDB()

	add := "/addseries"
	module := "LANTV"
	moduleType := "series"
	core.Home(w, add, module, moduleType)
}
func AddSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	back := "/tv"
	module := "LANTV"
	template := "addseries"
	submit := "/submitseries"
	route := "/updateSeriesSearch"
	previewRoute := "/populateSeries"
	core.Add(w, back, module, template, submit, route, previewRoute)
}

func SubmitSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	seriesID := r.FormValue("imdbCode")
	moduleType := "series"
	goTo := "/tv"
	core.SetMediaAdded(seriesID, moduleType)

	http.Redirect(w, r, goTo, http.StatusSeeOther)
}

func AddSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	AddSeason_(w, r, p)
}

func SubmitSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SubmitSeason_(w, r, p)
}

func SelectSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createSeriesDB()
	createEpisodesDB()

	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	//show the default series page
	series := p.ByName("seriesID")
	fmt.Printf("Playing episode: SeriesID=%s, SeasonNum=%d, EpisodeNum=%d\n", series, 1, 1)

	//need to mkae sure th "movieID" actually exists so it can 404 if it doesn't
	data, err := RetrieveSeriesFromDB(series)
	fmt.Printf("title is %s\n", data.Title)

	//if the series is not found, return a 404
	if err != nil {
		content.GenerateHTML(w, nil, "General", "notfound")
		return
	}

	episodes, _ := RetrieveEpisodesFromDB(data.Title)
	data = OrganizeIntoSeasons(data, episodes)

	//set the first season and episode to active
	if len(data.Seasons) > 0 {
		fmt.Printf("setting active\n")
		data.Seasons[0].Active = true
		data.Seasons[0].Episodes[0].Active = true
	}

	data.Back = "/tv"
	data.Add = ""

	data.Review = content.FormatParagraph(data.Review)

	SeriesData = Series{}
	SeriesData = data

	http.Redirect(w, r, "/tv/"+series, http.StatusSeeOther)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ShowSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	content.GenerateHTML(w, SeriesData, "LANTV", "series")
}

func SelectSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SelectSeason_(w, r, p)
}

func SelectEpisode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SelectEpisode_(w, r, p)
}

func UpdateSeriesSearch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	core.UpdateSearch(w, r, "tv")
}

func PopulateSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	PopulateSeries_(w, r, p)
}
