package lantv

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	db "webserver/src/DB"
)

func SetSeriesAdded(id string) {
	//update the added column in the series table
	_, err := db.Database.Exec("update series set added = 1 where id = ?", id)
	if err != nil {
		fmt.Printf("error adding series: %s\n", err)
	}
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

func createSeriesDB() {
	stmt, err := db.Database.Prepare(`
	CREATE TABLE IF NOT EXISTS series (
		uid INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT,
		added INTEGER DEFAULT 0,
		title TEXT,
		synopsis TEXT,
		release_date TEXT,
		runtime TEXT,
		seasons TEXT,
		rating TEXT,
		ratings TEXT,
		genres TEXT,
		cover_image TEXT,
		num_images INTEGER,
		review TEXT
	)
`)

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
	rows, err := db.Database.Query("select id, title, cover_image from series WHERE added = 1")
	for rows.Next() {
		se := Series{}
		if err = rows.Scan(&se.ID, &se.Title, &se.Image); err != nil {
			fmt.Printf("error getting series: %s\n", err)
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

func savePreviewToDB(key string, data []string) {
	// Convert genres to JSON string
	genresJSON, err := json.Marshal(data[7])
	if err != nil {
		fmt.Printf("error marshaling genres: %s\n", err)
		return
	}

	_, err = db.Database.Exec("insert into series (id, title, synopsis, release_date, runtime, seasons, rating, ratings, genres, cover_image, num_images, review) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		key, data[0], data[1], data[2], data[3], data[4], data[5], data[6], string(genresJSON), data[8], data[9], data[10])
	if err != nil {
		fmt.Printf("error adding series: %s\n", err)
	}
}

func retreivePreviewFromDB(id string) (data []string, success bool, err error) {
	rows, err := db.Database.Query("SELECT id, title, synopsis, release_date, runtime, seasons, rating, ratings, genres, cover_image, num_images, review FROM series WHERE id = ?", id)

	if err != nil {
		fmt.Printf("error retrieving series: %s\n", err)
		return data, success, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, synopsis, releaseDate, runtime, rating, seasons, ratings, genresJSON, coverImage, numImages, review string
		if err := rows.Scan(&id, &title, &synopsis, &releaseDate, &runtime, &seasons, &rating, &ratings, &genresJSON, &coverImage, &numImages, &review); err != nil {
			fmt.Printf("error scanning row: %s\n", err)
			return data, success, err
		}

		// Unmarshal genres JSON string
		var genres string
		if err := json.Unmarshal([]byte(genresJSON), &genres); err != nil {
			fmt.Printf("error unmarshaling genres: %s\n", err)
			success = false
			return data, success, err
		}

		genresStr := fmt.Sprintf("%v", genres)

		data = []string{title, synopsis, releaseDate, runtime, seasons, rating, ratings, genresStr, coverImage, numImages, review}
		success = true
	}

	if success {
		fmt.Printf("found series: %s\n", data[0])
	} else {
		fmt.Printf("series not found: %s\n", id)
	}

	return data, success, err
}

//return a list of every episode in the series in the db
func RetrieveSeriesFromDB(id string) (Series, error) {
	data, success, err := retreivePreviewFromDB(id)
	if !success {
		return Series{}, errors.New("series not found")
	}

	series := Series{}

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
