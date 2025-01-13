package core

import (
	"fmt"
	db "webserver/src/DB"
)

func SetMediaAdded(id string, moduleType string) {
	//update the added column in the series table
	query := fmt.Sprintf("UPDATE %s SET added = 1 WHERE id = ?", moduleType)
	_, err := db.Database.Exec(query, id)
	if err != nil {
		fmt.Printf("error adding series: %s\n", err)
	}
}

func getAllSeries(moduleType string) (series []Series, err error) {
	query := fmt.Sprintf("SELECT id, title, cover_image FROM %s WHERE added = 1", moduleType)
	rows, err := db.Database.Query(query)
	for rows.Next() {
		media := Series{}
		if err = rows.Scan(&media.ID, &media.Title, &media.Image); err != nil {
			fmt.Printf("error getting media: %s\n", err)
			return
		}
		series = append(series, media)
		fmt.Printf("id: %s, title: %s, image %s\n", media.ID, media.Title, media.Image)
	}
	if rows != nil {
		rows.Close()
	}

	fmt.Println("retrieving all")
	return
}

func AddPathToDB(moduleType string, path string, id string) bool {
	query := fmt.Sprintf("UPDATE %s SET path = ? WHERE id = ?", moduleType)
	_, err := db.Database.Exec(query, path, id)
	if err != nil {
		fmt.Printf("error adding path: %s\n", err)
		return false
	}

	return true
}
