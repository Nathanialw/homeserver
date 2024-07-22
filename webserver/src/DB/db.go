package db

import (
	"database/sql"
	"log"
)

type Movie struct {
	Uid         int
	Title       string
	Subtitle    string
	Director    string
	Cover       string
	ReleaseDate int
	genre       string
	Synopsis    string
}

var contentDB *sql.DB

func Init() {
	var err error
	contentDB, err = sql.Open("sqlite3", "../database/contentDB.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
}

func Movies(list string) (threads []Movie, err error) {
	var movies []Movie
	var movie Movie

	movies = append(movies, movie)

	return
}

// func Movies(list string) (threads []Movie, err error) {
// 	//get all the books from a readinglist
// 	//use the uid to get the books from the books table
// 	bookuids, err := contentDB.Query("select bookuid from readinglistbooks where reading_list = ?", list)

// 	//get list of uids
// 	//use the uids to get the books

// 	var rows *sql.Rows

// 	for bookuids.Next() {
// 		bookuids.Scan(&lst.Bookuid)
// 		fmt.Printf("bookuid: %d\n", lst.Bookuid)
// 		rows, err = contentDB.Query("select title, subtitle,  from books where uid = ?", lst.Bookuid)
// 		if err != nil {
// 			fmt.Printf("%s", err)
// 		}
// 		th := Movie{}
// 		for rows.Next() {
// 			if err = rows.Scan(&th.Title, &th.Subtitle); err != nil {
// 				fmt.Printf("%s", err)
// 				return
// 			}
// 			threads = append(threads, th)
// 			fmt.Printf("title: %s, author: %s\n", th.Title, th.Author)
// 		}
// 	}

// 	if rows != nil {
// 		rows.Close()
// 	}
// 	return
// }
