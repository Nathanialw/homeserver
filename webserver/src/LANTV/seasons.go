package lantv

import (
	"fmt"
	"net/http"
	"strconv"
	content "webserver/src/Content"
	core "webserver/src/Core"

	"github.com/julienschmidt/httprouter"
)

func SelectSeason_(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	series := r.FormValue("seriesID")
	currentSeason, _ := strconv.Atoi(r.FormValue("seasonNum"))
	episodeNum, _ := strconv.Atoi(r.FormValue("episodeNum"))

	fmt.Printf("Playing episode: SeriesID=%s, SeasonNum=%d, EpisodeNum=%d\n", series, currentSeason, episodeNum)

	var err error
	var data core.Series
	var episodes []core.Episode

	//need to ensure the "movieID" actually exists so it can 404 if it doesn't
	data, err = RetrieveSeriesFromDB(series)
	data.Synopsis = data.Title + `<br>` + data.Synopsis

	//organize the episodes by season
	episodes, _ = RetrieveEpisodesFromDB(data.Title)
	fmt.Printf("number of episodes: %d\n", len(episodes))

	data = OrganizeIntoSeasons(data, episodes)
	fmt.Printf("number of seasons: %d\n", len(data.Seasons))

	//reset the active season/episode
	for i := 0; i < len(data.Seasons); i++ {
		for j := 0; j < len(data.Seasons[i].Episodes); j++ {
			if i == currentSeason-1 && j == episodeNum-1 {
				data.Seasons[i].Episodes[j].Active = true
			} else {
				data.Seasons[i].Episodes[j].Active = false
			}
		}
		if i == currentSeason-1 {
			data.Seasons[i].Active = true
		} else {
			data.Seasons[i].Active = false
		}
	}

	data.Back = "/tv/"
	data.Add = "/addseason/" + series
	data.Review = content.FormatParagraph(data.Review)

	content.GenerateHTML(w, data, "LANTV", "series")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddSeason_(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	fmt.Printf("series title: %s\n", p.ByName("seriesID"))

	var data core.List
	// user.Session.LoggedIn = LoginStatus(r)
	// user.Session.Admin = AdminStatus(r)
	data.Back = "/tv/" + p.ByName("seriesID")
	data.Add = ""

	content.GenerateHTML(w, data, "LANTV", "addseason")
}

func SubmitSeason_(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	// fmt.Printf("series title: %s\n", p.ByName("seriesID"))

	// success, series := authenticateSeason(w, r)
	// if success {
	// 	insertSeasonIntoDB(series)
	// 	fmt.Print("added\n")
	// } else {
	// 	fmt.Print("not added\n")
	// }

	// http.Redirect(w, r, "/tv/"+p.ByName("seriesID"), http.StatusSeeOther)
}
