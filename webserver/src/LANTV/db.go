package lantv

import (
	"encoding/json"
	"fmt"
	"log"
	db "webserver/src/DB"
)

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
	stmt, err := db.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS episodes (
			uid INTEGER PRIMARY KEY AUTOINCREMENT, 
			seriesID TEXT, 
			episode TEXT, 
			season TEXT, 
			title TEXT, 
			subtitle TEXT, 
			image TEXT, 
			synopsis TEXT, 
			path TEXT
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

func RetrieveEpisodesFromDB(id string) ([]Episode, error) {
	var episodes []Episode
	rows, err := db.Database.Query("select * from episodes where seriesID = ?", id)
	if err != nil {
		fmt.Printf("error retrieving series: %s\n", err)
		return episodes, err
	}
	for rows.Next() {
		var episode Episode
		if err = rows.Scan(&episode.Uid, &episode.seriesID, &episode.EpisodeNum, &episode.Season, &episode.Title, &episode.Subtitle, &episode.Image, &episode.Synopsis, &episode.Path); err != nil {
			fmt.Printf("error scanning series: %s\n", err)
			return episodes, err
		}
		episodes = append(episodes, episode)
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Printf("number of episodes: %d\n", len(episodes))
	return episodes, nil
}

func OrganizeIntoSeasons(series Series, episodes []Episode) Series {
	// get number of seasons

	for i := range series.Seasons {
		series.Seasons[i].seriesID = series.ID
		series.Seasons[i].Active = false
		series.Seasons[i].SeasonNum = i + 1
	}

	numSeasons := 0
	for _, episode := range episodes {
		if episode.Season > numSeasons {
			numSeasons = episode.Season
		}
		series.Seasons[episode.Season-1].Episodes = append(series.Seasons[episode.Season-1].Episodes, episode)
	}

	//organize the episodes in each season
	for _, season := range series.Seasons {
		var organizedSeason = season

		for _, episode := range season.Episodes {
			organizedSeason.Episodes[episode.EpisodeNum-1] = episode
		}

		season = organizedSeason
	}

	//organize the episodes by season
	return series
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

func insertSeasonIntoDB(series Series) {
	for j := 0; j < len(series.Seasons); j++ {
		for i := 0; i < len(series.Seasons[j].Episodes); i++ {
			_, err := db.Database.Exec("insert into episodes (seriesID, episode, season, title, subtitle, image, synopsis, path) values (?, ?, ?, ?, ?, ?, ?, ?)", series.ID, i+1, series.Seasons[j].SeasonNum, series.Seasons[j].Episodes[i].Title, series.Seasons[j].Episodes[i].Subtitle, series.Seasons[j].Episodes[i].Image, series.Seasons[j].Episodes[i].Synopsis, series.Seasons[j].Episodes[i].Path)
			if err != nil {
				fmt.Printf("error adding series: %s\n", err)
			}
		}
	}
}
