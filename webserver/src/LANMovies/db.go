package lanmovies

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
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
			review TEXT,
			path TEXT DEFAULT " "
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
	rows, err := db.Database.Query("SELECT id, title, synopsis, release_date, runtime, rating, ratings, genres, cover_image, num_images, review, path FROM movies WHERE id = ?", id)

	if err != nil {
		fmt.Printf("retreivePreviewFromDB error retrieving movies: %s\n", err)
		return data, success, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, title, synopsis, releaseDate, runtime, rating, ratings, genresJSON, coverImage, numImages, review, path string
		if err := rows.Scan(&id, &title, &synopsis, &releaseDate, &runtime, &rating, &ratings, &genresJSON, &coverImage, &numImages, &review, &path); err != nil {
			fmt.Printf("retreivePreviewFromDB error scanning row: %s\n", err)
			return data, success, err
		}

		// Unmarshal genres JSON string
		var genres string
		if err := json.Unmarshal([]byte(genresJSON), &genres); err != nil {
			fmt.Printf("retreivePreviewFromDB error unmarshaling genres: %s\n", err)
			success = false
			return data, success, err
		}

		genresStr := fmt.Sprintf("%v", genres)

		data = []string{title, synopsis, releaseDate, runtime, "", rating, ratings, genresStr, coverImage, numImages, review, path, " "}
		success = true
	}

	if success {
		fmt.Printf("retreivePreviewFromDB found movies: %s\n", data[0])
	} else {
		fmt.Printf("retreivePreviewFromDB movies not found: %s\n", id)
	}

	return data, success, err
}

func savePreviewToDB(key string, data []string) {
	//if id is in db
	rows, err := db.Database.Query("SELECT id FROM movies WHERE id = ?", key)
	if err != nil {
		fmt.Printf("savePreviewToDB movie already in db: %s not adding another\n", key)
		return
	}
	defer rows.Close()
	if rows.Next() {
		fmt.Printf("savePreviewToDB movie already in db: %s not adding another\n", key)
		return
	}

	// Convert genres to JSON string
	genresJSON, err := json.Marshal(data[7])
	if err != nil {
		fmt.Printf("savePreviewToDB error marshaling genres: %s\n", err)
		return
	}

	_, err = db.Database.Exec("insert into movies (id, title, synopsis, release_date, runtime, rating, ratings, genres, cover_image, num_images, review, path) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		key, data[0], data[1], data[2], data[3], data[5], data[6], string(genresJSON), data[8], data[9], data[10], " ")

	if err != nil {
		fmt.Printf("savePreviewToDB error adding movies: %s\n", err)
	}
}

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
	movie.Rating = data[5]
	movie.Ratings = data[6]
	movie.GenresList = strings.Split(data[7], ", ")
	movie.Image = data[8]
	movie.NumImages = data[9]
	movie.Review = data[10]
	movie.Path = data[11]

	return movie, err
}
