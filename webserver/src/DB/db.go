package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Movie struct {
	Uid      string
	Title    string
	Subtitle string
	Director string
	Cover    string
	Year     int
	Genre    string
	Synopsis string
}

var Database *sql.DB

func Init() {
	var err error
	//if the db doesn't exist, create it in the correct place

	Database, err = sql.Open("sqlite3", "../db/homeserver.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
}

func close() {
	Database.Close()
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
