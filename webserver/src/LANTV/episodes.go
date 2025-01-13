package lantv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	content "webserver/src/Content"
	db "webserver/src/DB"

	"github.com/julienschmidt/httprouter"
)

type VideoResponse struct {
	VideoURL string `json:"videoURL"`
}

func SelectEpisode_(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	seriesTitle := r.FormValue("seriesID")
	currentSeason, _ := strconv.Atoi(r.FormValue("seasonNum"))
	episodeNum, _ := strconv.Atoi(r.FormValue("episodeNum"))

	fmt.Printf("Playing episode: SeriesID=%s, SeasonNum=%d, EpisodeNum=%d\n", seriesTitle, currentSeason, episodeNum)

	var episodes []Episode

	//need to ensure the "movieID" actually exists so it can 404 if it doesn't
	data, err := RetrieveSeriesFromDB(seriesTitle)

	//organize the episodes by season
	episodes, _ = RetrieveEpisodesFromDB(data.ID)
	fmt.Printf("number of episodes: %d\n", len(episodes))

	data = OrganizeIntoSeasons(data, episodes)
	fmt.Printf("number of seasons: %d\n", len(data.Seasons))

	data.Review = content.FormatParagraph(data.Review)

	videoURL := data.Seasons[currentSeason-1].Episodes[episodeNum-1].Path

	response := VideoResponse{VideoURL: videoURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAllEpisodes(series string) (episodes []Episode, err error) {
	rows, err := db.Database.Query("select title, subtitle, image from episodes where series = ?", series)
	for rows.Next() {
		ep := Episode{}
		if err = rows.Scan(&ep.Title, &ep.Subtitle, &ep.Image); err != nil {
			fmt.Printf("%s", err)
			return
		}
		episodes = append(episodes, ep)
		fmt.Printf("name: %s\n", ep.Title)
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Println("retrieving all")
	return
}
