package lanbooks

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

type PageData struct {
	Title string
	Body  string
}

func Home (w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "LAN Books",
		Body:  "Welcome to LAN Books",
	}
	tmpl, err := template.ParseFiles("templates/LANBooks.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}