package lanmovies

import (
	"encoding/json"
	"fmt"
	"net/http"
	content "webserver/src/Content"
	core "webserver/src/Core"

	"github.com/julienschmidt/httprouter"
)

func PopulateMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := r.FormValue("id")

	results, success, _ := retreivePreviewFromDB(id) //returns wihtout season
	if !success {
		fmt.Printf("PopulateMovie Scraping movie data: %s\n", id)
		results = core.Preview_Series(id) //returns wiht season
		savePreviewToDB(id, results)
	}

	results[len(results)-1] = content.FormatParagraph(results[len(results)-1])

	response, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("PopulateMovie error: %s\n", err)
		http.Error(w, "PopulateMovie Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
