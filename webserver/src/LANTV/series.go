package lantv

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	content "webserver/src/Content"
	core "webserver/src/Core"

	"github.com/julienschmidt/httprouter"
)

func PopulateSeries_(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	id := r.FormValue("id")

	results, success, _ := retreivePreviewFromDB(id)
	if !success {
		results = core.Preview_Series(id)
		savePreviewToDB(id, results)
	}
	results[10] = content.FormatParagraph(results[10])

	response, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

//return a list of every episode in the series in the db
func RetrieveSeriesFromDB(id string) (core.Series, error) {
	data, success, err := retreivePreviewFromDB(id)
	if !success {
		return core.Series{}, errors.New("series not found")
	}

	series := core.Series{}

	series.ID = id
	series.Title = data[0]
	series.Synopsis = data[1]
	series.ReleaseDate = data[2]
	series.Runtime = data[3]
	series.NumSeasons = data[4]
	series.Rating = data[5]
	series.Ratings = data[6]
	series.Genres = data[7]
	series.Image = data[8]
	series.NumImages = data[9]
	series.Review = data[10]

	return series, err
}
