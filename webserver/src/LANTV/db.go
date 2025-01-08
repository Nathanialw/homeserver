package lantv

import (
	"fmt"
	"log"
	db "webserver/src/DB"
)

func insertSeriesIntoDB(series Series) {
	//insert series
	_, err := db.Database.Exec("insert into series (title, subtitle, image, synopsis, path) values (?, ?, ?, ?, ?)", series.Title, series.Subtitle, series.Image, series.Synopsis, series.Path)
	if err != nil {
		fmt.Printf("error adding series: %s\n", err)
	}

	//insert episodes
}

func insertSeasonIntoDB(series Series) {
	for j := 0; j < len(series.Seasons); j++ {
		for i := 0; i < len(series.Seasons[j].Episodes); i++ {
			_, err := db.Database.Exec("insert into episodes (series, episode, season, title, subtitle, image, synopsis, path) values (?, ?, ?, ?, ?, ?, ?, ?)", currentSeriesTitle, i+1, series.Seasons[j].SeasonNum, series.Seasons[j].Episodes[i].Title, series.Seasons[j].Episodes[i].Subtitle, series.Seasons[j].Episodes[i].Image, series.Seasons[j].Episodes[i].Synopsis, series.Seasons[j].Episodes[i].Path)
			if err != nil {
				fmt.Printf("error adding series: %s\n", err)
			}
		}
	}
}

func RetrieveEpisodesFromDB(title string) ([]Episode, error) {
	var episodes []Episode
	rows, err := db.Database.Query("select * from episodes where series = ?", title)
	if err != nil {
		fmt.Printf("error retrieving series: %s\n", err)
		return episodes, err
	}
	for rows.Next() {
		var episode Episode
		if err = rows.Scan(&episode.Uid, &episode.Series, &episode.EpisodeNum, &episode.Season, &episode.Title, &episode.Subtitle, &episode.Image, &episode.Synopsis, &episode.Path); err != nil {
			fmt.Printf("error scanning series: %s\n", err)
			return episodes, err
		}
		episodes = append(episodes, episode)
	}
	if rows != nil {
		rows.Close()
	}

	return episodes, nil
}

//return a list of every episode in the series in the db
func RetrieveSeriesFromDB(title string) (Series, error) {
	var series Series
	rows, err := db.Database.Query("select * from series where title = ?", title)
	if err != nil {
		fmt.Printf("error retrieving series: %s\n", err)
		return series, err
	}
	for rows.Next() {
		if err = rows.Scan(&series.Uid, &series.Title, &series.Subtitle, &series.Image, &series.Synopsis, &series.Path); err != nil {
			fmt.Printf("error scanning series: %s\n", err)
			return series, err
		}
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Printf("image: %s, path: %s\n", series.Image, series.Path)
	return series, nil
}

func createSeriesDB() {
	stmt, err := db.Database.Prepare("CREATE TABLE IF NOT EXISTS series (uid INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, subtitle TEXT, image TEXT, synopsis TEXT, path TEXT)")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v\n", err)
	}
	_, execErr := stmt.Exec()
	if execErr != nil {
		log.Fatalf("Failed to execute table creation statement: %v", execErr)
	}
}

func createEpisodesDB() {
	stmt, err := db.Database.Prepare("CREATE TABLE IF NOT EXISTS episodes (uid INTEGER PRIMARY KEY AUTOINCREMENT, series TEXT, episode TEXT, season TEXT, title TEXT, subtitle TEXT, image TEXT, synopsis TEXT, path TEXT)")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v\n", err)
	}
	_, execErr := stmt.Exec()
	if execErr != nil {
		log.Fatalf("Failed to execute table creation statement: %v", execErr)
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

func getAllSeasons(series string) (seasons []Season, err error) {
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

func getAllSeries() (series []Series, err error) {
	rows, err := db.Database.Query("select title, subtitle, image from series")
	for rows.Next() {
		se := Series{}
		if err = rows.Scan(&se.Title, &se.Subtitle, &se.Image); err != nil {
			fmt.Printf("%s", err)
			return
		}
		series = append(series, se)
		fmt.Printf("name: %s\n", se.Title)
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Println("retrieving all")
	return
}

func getSeries() (series []Series, err error) {
	rows, err := db.Database.Query("select title, subtitle, image from series")
	for rows.Next() {
		se := Series{}
		if err = rows.Scan(&se.Title, &se.Subtitle, &se.Image); err != nil {
			fmt.Printf("%s", err)
			return
		}
		series = append(series, se)
		fmt.Printf("name: %s\n", se.Title)
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Println("retrieving all")
	return
}
