package lanmovies

import (
	"fmt"
	"log"
	"net/http"
	content "webserver/src/Content"
	db "webserver/src/DB"

	"github.com/julienschmidt/httprouter"
)

type Movie struct {
	Uid      string
	Title    string
	Subtitle string
	Director string
	Cover    string
	Year     int
	Length   int
	Genre    string
	Synopsis string
}

type PageData struct {
	Title string
	Body  string
}

type UserSession struct {
	Movies []Movie

	//post history
	//post history json
}

func createMoviesDB() {
	stmt, err := db.Database.Prepare("CREATE TABLE IF NOT EXISTS movies (uid INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, subtitle TEXT, director TEXT, cover TEXT, year INTEGER, length INTEGER, genre TEXT, synopsis TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

func getAll() (threads []Movie, err error) {
	rows, err := db.Database.Query("select title, director, cover, synopsis from movies")
	for rows.Next() {
		th := Movie{}
		if err = rows.Scan(&th.Title, &th.Director, &th.Cover, &th.Synopsis); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.Title, th.Genre)
	}
	if rows != nil {
		rows.Close()
	}
	fmt.Printf("retrieving all")
	return
}

func getGenre(genre string) (threads []Movie, err error) {
	rows, err := db.Database.Query("select title, director, cover, synopsis from movies where genre = ?", genre)
	for rows.Next() {
		th := Movie{}
		if err = rows.Scan(&th.Title, &th.Director, &th.Cover, &th.Synopsis); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.Title, th.Genre)
	}
	if rows != nil {
		rows.Close()
	}
	fmt.Printf("retrieving genre %s", genre)
	return
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createMoviesDB()

	data, err := getAll()
	content.GenerateHTML(w, data, "LANMovies")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func unused_movied(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// currentPage = r.URL.Path

	// fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	//strip off the end of the url
	// fmt.Printf("category: %s\n", list)
	// data.LoggedIn = LoginStatus(r)
	// data.Admin = AdminStatus(r)
	// data.Category, _ = GetCategory(list)
	// if data.Category.Category == "" {
	// 	notfound(w, r, p)
	// 	return
	// }

	// content.GenerateHTML(w, data, "LANMovies")
}
