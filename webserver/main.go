package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	lanbooks "webserver/modules/LANBooks"
	landocs "webserver/modules/LANDocs"
	langames "webserver/modules/LANGames"
	lanmovies "webserver/modules/LANMovies"
	lanmusic "webserver/modules/LANMusic"
	lanpics "webserver/modules/LANPics"
	lantv "webserver/modules/LANTV"
)

type PageData struct {
	Title string
	Body  string
}

func main() {
	r := httprouter.New()
	
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
		Title: "My Page Title",
		Body:  "Welcome to my dwebsite!",
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