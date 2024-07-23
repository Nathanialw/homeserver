package lanbooks

import (
	"fmt"
	"net/http"
	content "webserver/src/Content"

	"github.com/julienschmidt/httprouter"
)

type PageData struct {
	Title string
	Body  string
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	data := PageData{
		Title: "LAN Books",
		Body:  "Welcome to LAN Books",
	}
	content.GenerateHTML(w, data, "LANBooks", "LANBooks")
	// tmpl, err := template.ParseFiles("../modules/LANBooks/templates/LANBooks.html")
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
