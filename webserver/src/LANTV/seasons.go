package lantv

import (
	"fmt"
	"net/http"
	"strconv"
	content "webserver/src/Content"
	core "webserver/src/Core"
	db "webserver/src/DB"

	"github.com/julienschmidt/httprouter"
)

func GetSeason(w http.ResponseWriter, r *http.Request, p httprouter.Params) (Series, error) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	series := p.ByName("seriesID")
	currentSeason, _ := strconv.Atoi(p.ByName("seasonNum"))

	fmt.Printf("Playing episode: SeriesID=%s, SeasonNum=%d, EpisodeNum=%d\n", series, currentSeason, 1)

	//need to ensure the "movieID" actually exists so it can 404 if it doesn't
	data, err := RetrieveSeriesFromDB(series)

	//organize the episodes by season
	episodes, _ := RetrieveEpisodesFromDB(data.Title)
	data = OrganizeIntoSeasons(data, episodes)

	if len(data.Seasons) > 0 && currentSeason > 0 {
		data.Seasons[currentSeason-1].Active = true
		if len(data.Seasons[currentSeason-1].Episodes) > 0 {
			data.Seasons[currentSeason-1].Episodes[0].Active = true
		}
	}

	data.Back = "/tv/"
	data.Add = ""
	data.Review = content.FormatParagraph(data.Review)

	return data, err
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

func getAllSeasons(series string, moduleType string) (seasons []Season, err error) {
	rows, err := db.Database.Query("select season from season where series = ?", series)
	for rows.Next() {
		se := Season{}
		if err = rows.Scan(&se.Image); err != nil {
			fmt.Printf("%s", err)
			return
		}
		seasons = append(seasons, se)
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Println("retrieving all")
	return
}
