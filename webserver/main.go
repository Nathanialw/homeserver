package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	db "webserver/src/DB"
	lanbooks "webserver/src/LANBooks"
	landocs "webserver/src/LANDocs"
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

type UserSession struct {
	Username string
	Admin    bool
	LoggedIn bool
}

func main() {
	db.Init()
	r := httprouter.New()

	fs := http.FileServer(http.Dir("../../public/"))
	r.NotFound = http.StripPrefix("/", fileServerWith404(fs))

	r.GET("/", home)
	r.GET("/books", lanbooks.Home)
	r.GET("/docs", landocs.Home)
	r.GET("/games", langames.Home)
	r.GET("/movies", lanmovies.Home)
	r.GET("/music", lanmusic.Home)
	r.GET("/pics", lanpics.Home)
	r.GET("/tv", lantv.Home)

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
	tmpl, err := template.ParseFiles("../templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	var data UserSession
	//data.LoggedIn = LoginStatus(r)
	tmpl, err := template.ParseFiles("../templates/notfound.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
