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
	seriesID   string
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
	seriesID  string
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

func SubmitSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SubmitSeason_(w, r, p)
}

func ShowSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, err := GetSeason(w, r, p)
	content.GenerateHTML(w, data, "LANTV", "series", "aboutmedia")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func ShowSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	content.GenerateHTML(w, SeriesData, "LANTV", "series", "aboutmedia")
// }

func SelectEpisode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SelectEpisode_(w, r, p)
}

func UpdateSeriesSearch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	core.UpdateSearch(w, r, "tv")
}

func PopulateSeries(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	PopulateSeries_(w, r, p)
}

func SubmitSeasonFolder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("seriesID")
	seasonNum := p.ByName("seasonNum")

	route := "/tv/" + id + "/" + seasonNum
	fmt.Printf("message received from %s, series: %s, season: %s, route: %s\n", r.RemoteAddr, id, seasonNum, route)

	media := r.FormValue("media")
	fmt.Printf("media: %s\n", media)

	SubmitSeason_(w, r, p)
	http.Redirect(w, r, route, http.StatusSeeOther)
}
