package lanmovies

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	content "webserver/src/Content"
	core "webserver/src/Core"

	"github.com/julienschmidt/httprouter"
)

//return a list of every episode in the series in the db
func RetrieveMovieFromDB(id string) (Movie, error) {
	data, success, err := retreivePreviewFromDB(id)
	if !success {
		return Movie{}, errors.New("series not found")
	}

	movie := Movie{}

	movie.ID = id
	movie.Title = data[0]
	movie.Synopsis = data[1]
	movie.ReleaseDate = data[2]
	movie.Runtime = data[3]
	movie.Rating = data[4]
	movie.Ratings = data[5]
	movie.Genres = strings.Split(data[6], ", ")
	movie.Image = data[7]
	movie.NumImages = data[8]
	movie.Review = data[9]
	movie.Path = data[10]

	return movie, err
}

func PopulateMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	id := r.FormValue("id")

	results, success, _ := retreivePreviewFromDB(id)
	if !success {
		results = core.Preview_Series(id)
		savePreviewToDB(id, results)
	}

	results[len(results)-1] = content.FormatParagraph(results[len(results)-1])

	response, err := json.Marshal(results)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
