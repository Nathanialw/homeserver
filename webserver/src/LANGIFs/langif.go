package langif

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
		Title: "LAN GIF",
		Body:  "Welcome to LAN GIF",
	}

	content.GenerateHTML(w, data, "LANGIF", "LANGIF")
}
