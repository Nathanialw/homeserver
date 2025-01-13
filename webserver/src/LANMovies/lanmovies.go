package lanmovies

import (
	"fmt"
	"net/http"
	"strings"
	content "webserver/src/Content"
	core "webserver/src/Core"

	"github.com/julienschmidt/httprouter"
)

type Movie struct {
	ID          string
	Title       string
	Synopsis    string
	ReleaseDate string
	Runtime     string
	Rating      string
	Ratings     string
	Genres      []string
	Image       string
	NumImages   string
	Review      string
	Path        string

	Back string
	Add  string

	Director string
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createMoviesDB()
	add := "/addmovie"
	module := "LANMovies"
	moduleType := "movies"

	core.Home(w, add, module, moduleType)
}

func AddMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	back := "/movie"
	module := "LANMovies"
	template := "addmovie"
	submit := "/submitmovie"
	route := "/updateMovieSearch"
	previewRoute := "/populateMovie"

	core.Add(w, back, module, template, submit, route, previewRoute)
}

func SubmitMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	seriesID := r.FormValue("imdbCode")
	moduleType := "movies"
	home := "/movie"
	core.SetMediaAdded(seriesID, moduleType)

	http.Redirect(w, r, home, http.StatusSeeOther)
}

func SubmitMovieFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	refresh := strings.Split(r.URL.Path, "/")
	id := refresh[2]

	moduleType := "movies"
	pathSuffix := "/movie/" + id

	if Authenticate() {
		path := core.SubmitFile(w, r, pathSuffix)
		core.AddPathToDB(moduleType, path, id)
	}
	http.Redirect(w, r, pathSuffix, http.StatusSeeOther)
}

func ShowMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createMoviesDB()

	//need to mkae sure th "movieID" actually exists so it can 404 if it doesn't
	data, err := RetrieveMovieFromDB(p.ByName("movieID"))
	fmt.Printf("title is %s\n", data.Title)

	//if the series is not found, return a 404
	if err != nil {
		content.GenerateHTML(w, nil, "General", "notfound")
		return
	}

	data.Back = "/movie"
	data.Add = ""

	data.Review = content.FormatParagraph(data.Review)

	content.GenerateHTML(w, data, "LANMovies", "movie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateMoviesSearch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	core.UpdateSearch(w, r, "movies")
}
