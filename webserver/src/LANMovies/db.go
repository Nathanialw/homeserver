package lanmovies

import (
	"encoding/json"
	"fmt"
	"log"
	db "webserver/src/DB"
)

func createMoviesDB() {
	stmt, err := db.Database.Prepare(`
		CREATE TABLE IF NOT EXISTS movies (
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			id TEXT,
			added INTEGER DEFAULT 0,
			title TEXT,
			synopsis TEXT,
			release_date TEXT,
			runtime TEXT,
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

func retreivePreviewFromDB(id string) (data []string, success bool, err error) {
	rows, err := db.Database.Query("SELECT id, title, synopsis, release_date, runtime, rating, ratings, genres, cover_image, num_images, review FROM movies WHERE id = ?", id)

	if err != nil {
		fmt.Printf("error retrieving movies: %s\n", err)
		return data, success, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, synopsis, releaseDate, runtime, rating, ratings, genresJSON, coverImage, numImages, review string
		if err := rows.Scan(&id, &title, &synopsis, &releaseDate, &runtime, &rating, &ratings, &genresJSON, &coverImage, &numImages, &review); err != nil {
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

		data = []string{title, synopsis, releaseDate, runtime, rating, ratings, genresStr, coverImage, numImages, review}
		success = true
	}

	if success {
		fmt.Printf("found movies: %s\n", data[0])
	} else {
		fmt.Printf("movies not found: %s\n", id)
	}

	return data, success, err
}

func savePreviewToDB(key string, data []string) {
	// Convert genres to JSON string
	genresJSON, err := json.Marshal(data[6])
	if err != nil {
		fmt.Printf("error marshaling genres: %s\n", err)
		return
	}

	_, err = db.Database.Exec("insert into movies (id, title, synopsis, release_date, runtime, rating, ratings, genres, cover_image, num_images, review) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		key, data[0], data[1], data[2], data[3], data[4], data[5], string(genresJSON), data[7], data[8], data[9])
	if err != nil {
		fmt.Printf("error adding movies: %s\n", err)
	}
}
