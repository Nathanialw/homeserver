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
