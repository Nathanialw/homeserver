package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	content "webserver/src/Content"
	db "webserver/src/DB"
	lanbooks "webserver/src/LANBooks"
	langames "webserver/src/LANGames"
	lanmovies "webserver/src/LANMovies"
	lanmusic "webserver/src/LANMusic"
	lanpics "webserver/src/LANPics"
	lantv "webserver/src/LANTV"

	"github.com/julienschmidt/httprouter"
)

type PageData struct {
	Title string
	Body  string
	Back  string
	Add   string
}

func StripPrefix(w http.ResponseWriter, r *http.Request, prefix string, toStrip string) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	if _, err := os.Stat(toStrip + filePath); os.IsNotExist(err) {
		notfound(w, r, httprouter.Params{})
		return
	}
	http.StripPrefix(prefix, http.FileServer(http.Dir(toStrip))).ServeHTTP(w, r)
}

func main() {
	db.Init()

	r := httprouter.New()

	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve static files from /mnt/media, strip the leading part of the URL path
		fmt.Printf("serving %s\n", r.URL.Path)
		if strings.HasPrefix(r.URL.Path, "/mnt/media/") {
			http.StripPrefix("/mnt/media/", http.FileServer(http.Dir("/mnt/media"))).ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/movie/") {
			StripPrefix(w, r, "/movie/", "../../public/")
		} else if strings.HasPrefix(r.URL.Path, "/tv/") {
			StripPrefix(w, r, "/tv/", "../../public/")
		} else if strings.HasPrefix(r.URL.Path, "/books/") {
			StripPrefix(w, r, "/books/", "../../public/")
		} else if strings.HasPrefix(r.URL.Path, "/books/") {
			StripPrefix(w, r, "/music/", "../../public/")
		} else if strings.HasPrefix(r.URL.Path, "/books/") {
			StripPrefix(w, r, "/games/", "../../public/")
		} else if strings.HasPrefix(r.URL.Path, "/books/") {
			StripPrefix(w, r, "/pics/", "../../public/")
		} else {
			StripPrefix(w, r, "/", "../../public/")
		}
	})

	r.GET("/", home)

	r.GET("/books", lanbooks.Home)
	r.GET("/book/:bookID", lanbooks.Show)
	r.GET("/addbook", lanbooks.AddBook)
	r.POST("/submitbook", lanbooks.SubmitBook)

	r.GET("/movie", lanmovies.Home)
	r.GET("/movie/:movieID", lanmovies.ShowMovie)
	r.GET("/addmovie", lanmovies.AddMovie)
	r.POST("/submitmovie", lanmovies.SubmitMovie)
	r.POST("/removemovie", lanmovies.RemoveMovie)

	r.GET("/tv", lantv.Home)
	r.GET("/addseries", lantv.AddSeries)
	r.POST("/submitseries", lantv.SubmitSeries)
	r.GET("/addseason/:seriesID", lantv.AddSeason)
	r.POST("/submitseason", lantv.SubmitSeason)
	r.GET("/tv/:seriesID", lantv.ShowSeries)
	r.POST("/selectSeason", lantv.SelectSeason)
	r.POST("/selectEpisode", lantv.SelectEpisode)

	r.GET("/games", langames.Home)
	r.GET("/music", lanmusic.Home)
	r.GET("/pics", lanpics.Home)

	address := "127.0.0.1:"
	if len(os.Args) > 1 {
		address += os.Args[1]
	} else {
		address += "10002"
	}

	server := http.Server{
		Addr:    address,
		Handler: r,
	}

	fmt.Println("Running at address: ", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{}

	content.GenerateHTML(w, data, "General", "home")
}

func notfound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Printf("not found received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{}

	content.GenerateHTML(w, data, "General", "notfound")
}
