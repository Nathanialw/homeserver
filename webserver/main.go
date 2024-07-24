package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	content "webserver/src/Content"
	db "webserver/src/DB"
	lanbooks "webserver/src/LANBooks"
	landocs "webserver/src/LANDocs"
	langif "webserver/src/LANGIFs"
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
}

func main() {
	db.Init()
	r := httprouter.New()

	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Serve static files from /mnt/media, strip the leading part of the URL path
		if strings.HasPrefix(r.URL.Path, "/mnt/media/") {
			http.StripPrefix("/mnt/media/", http.FileServer(http.Dir("/mnt/media"))).ServeHTTP(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/movie/") {
			http.StripPrefix("/movie/", http.FileServer(http.Dir("../../public/"))).ServeHTTP(w, r)
			notfound(w, r, httprouter.Params{})
		} else {
			http.StripPrefix("/", http.FileServer(http.Dir("../../public/"))).ServeHTTP(w, r)
			notfound(w, r, httprouter.Params{})
		}
	})

	r.GET("/", home)
	r.GET("/books", lanbooks.Home)
	r.GET("/docs", landocs.Home)
	r.GET("/games", langames.Home)

	r.GET("/movies", lanmovies.Home)
	r.GET("/movie/:movieID", lanmovies.ShowMovie)
	r.GET("/addmovie", lanmovies.AddMovie)
	r.POST("/submitmovie", lanmovies.SubmitMovie)

	r.GET("/music", lanmusic.Home)
	r.GET("/pics", lanpics.Home)
	r.GET("/tv", lantv.Home)
	r.GET("/gifs", langif.Home)

	server := http.Server{
		Addr:    "127.0.0.1:10002",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}

func home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "Landing Page",
		Body:  "Welcome to the home server landing page",
	}

	content.GenerateHTML(w, data, "Content", "home")
}

func fileServerWith404(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := os.Stat("../../public/" + r.URL.Path)

		if os.IsNotExist(err) {
			fmt.Printf("file does not exist: %s\n", r.URL.Path)
			// If the file does not exist, serve your 404 page
			notfound(w, r, httprouter.Params{})
			return
		}

		fmt.Printf("file exists: %s\n", r.URL.Path)
		// If the file exists, serve it
		h.ServeHTTP(w, r)
	}
}

func notfound(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{}

	content.GenerateHTML(w, data, "Content", "notfound")
}
