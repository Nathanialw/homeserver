package lanmovies

import (
	"fmt"
	"log"
	"net/http"
	authenticate "webserver/src/Authenticate"
	content "webserver/src/Content"
	db "webserver/src/DB"
	user "webserver/src/User"

	"github.com/julienschmidt/httprouter"
)

type Movie struct {
	Uid      string
	Title    string
	Subtitle string
	Director string
	Image    string
	Year     int
	Length   int
	Genre    string
	Synopsis string
}

type MoviesList struct {
	User     user.Session
	NotEmpty bool
	Movies   []Movie
}

func createMoviesDB() {
	stmt, err := db.Database.Prepare("CREATE TABLE IF NOT EXISTS movies (uid INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, subtitle TEXT, director TEXT, image TEXT, year INTEGER, length INTEGER, genre TEXT, synopsis TEXT, series TEXT, path TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

// func getAll() (threads []Movie, err error) {
func getAll() (threads []Movie, err error) {
	rows, err := db.Database.Query("select title, director, image, synopsis from movies")
	for rows.Next() {
		th := Movie{}
		if err = rows.Scan(&th.Title, &th.Director, &th.Image, &th.Synopsis); err != nil {
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

func getByGenre(filter string) (threads []Movie, err error) {
	rows, err := db.Database.Query("select title, director, image, synopsis from movies where genre = ?", filter)
	if err != nil {
		fmt.Printf("query failed: %s\n", err)
	}

	for rows.Next() {
		th := Movie{}
		if err = rows.Scan(&th.Title, &th.Director, &th.Image, &th.Synopsis); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.Title, th.Genre)
	}
	if rows != nil {
		rows.Close()
	}

	return
}

func getByDirector(filter string) (threads []Movie, err error) {
	rows, err := db.Database.Query("select title, director, image, synopsis from movies where director = ?", filter)
	if err != nil {
		fmt.Printf("query failed: %s\n", err)
	}

	for rows.Next() {
		th := Movie{}
		if err = rows.Scan(&th.Title, &th.Director, &th.Image, &th.Synopsis); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.Title, th.Genre)
	}
	if rows != nil {
		rows.Close()
	}

	return
}

func getBySeries(filter string) (threads []Movie, err error) {
	rows, err := db.Database.Query("select title, director, image, synopsis from movies where series = ?", filter)
	if err != nil {
		fmt.Printf("query failed: %s\n", err)
	}

	for rows.Next() {
		th := Movie{}
		if err = rows.Scan(&th.Title, &th.Director, &th.Image, &th.Synopsis); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s, category: %s\n", th.Title, th.Genre)
	}
	if rows != nil {
		rows.Close()
	}

	return
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createMoviesDB()

	var err error
	var data MoviesList

	// data.Movies, err = getAll()
	data.Movies, err = getAll()
	if len(data.Movies) > 0 {
		data.NotEmpty = true
		fmt.Printf("The director is %s.\n", data.Movies[0].Director)
	} else {
		data.NotEmpty = false
		fmt.Printf("none found\n")
	}
	content.GenerateHTML(w, data, "LANMovies", "LANMovies")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//func unused_movied(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
// }

func AuthenticateMovie(w http.ResponseWriter, r *http.Request) bool {
	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	director := r.FormValue("director")
	year := r.FormValue("year")
	series := r.FormValue("series")
	length := r.FormValue("length")
	synopsis := r.FormValue("synopsis")

	imageFile, imageFilename := authenticate.FormVideo("image", r)
	videoFile, videoFilename := authenticate.FormVideo("media", r)

	imageFile.Close()
	videoFile.Close()
	fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s, %s\n", title, subtitle, director, imageFilename, videoFilename, year, series, length, synopsis)

	return authenticate.ValidText(title) ||
		authenticate.ValidText(subtitle) ||
		authenticate.ValidText(director) ||
		authenticate.ValidYear(year) ||
		authenticate.ValidLength(length) ||
		authenticate.ValidText(series) ||
		authenticate.ValidText(synopsis)
	// authenticate.ValidImage()
	// authenticate.ValidVideo()
}

func AddMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data MoviesList
	// user.Session.LoggedIn = LoginStatus(r)
	// user.Session.Admin = AdminStatus(r)
	content.GenerateHTML(w, data, "LANMovies", "addmovie")
}

func SubmitMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	//check on the browser side first and give feedback to the user
	//then check on the server side
	//if this fails it should redirect to the addbook page

	// var data user.Session
	// user.Session.LoggedIn = LoginStatus(r)

	// if !data.LoggedIn {
	// 	fmt.Println("not logged in")
	// 	notfound(w, r, p)
	// 	return
	// }

	// VerifyAndInsertBook(w, r, contentDB)
	AuthenticateMovie(w, r)
	fmt.Print("added\n")
	http.Redirect(w, r, "/addmovie", http.StatusSeeOther)
}
