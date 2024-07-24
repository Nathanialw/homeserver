package lanmovies

import (
	"fmt"
	"log"
	"net/http"
	authenticate "webserver/src/Authenticate"
	content "webserver/src/Content"
	db "webserver/src/DB"
	upload "webserver/src/Upload"
	user "webserver/src/User"

	"github.com/julienschmidt/httprouter"
)

type Movie struct {
	Uid      string
	Title    string
	Subtitle string
	Director string
	Image    string
	Year     string
	Length   string
	Genre    string
	Series   string
	Synopsis string
	Path     string
}

type MoviesList struct {
	User     user.Session
	NotEmpty bool
	Movies   []Movie
}

func createMoviesDB() {
	stmt, err := db.Database.Prepare("CREATE TABLE IF NOT EXISTS movies (uid INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, subtitle TEXT, director TEXT, image TEXT, year INTEGER, length INTEGER, genre TEXT, synopsis TEXT, series TEXT, path TEXT)")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v\n", err)
	}
	_, execErr := stmt.Exec()
	if execErr != nil {
		log.Fatalf("Failed to execute table creation statement: %v", execErr)
	}
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

func ShowMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createMoviesDB()

	var err error
	var data Movie
	data, err = RetrieveMovieFromDB("c")

	content.GenerateHTML(w, data, "LANMovies", "movie")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RetrieveMovieFromDB(title string) (Movie, error) {
	var movie Movie
	rows, err := db.Database.Query("select * from movies where title = ?", title)
	if err != nil {
		fmt.Printf("error retrieving movie: %s\n", err)
		return movie, err
	}
	for rows.Next() {
		if err = rows.Scan(&movie.Uid, &movie.Title, &movie.Subtitle, &movie.Director, &movie.Image, &movie.Year, &movie.Length, &movie.Genre, &movie.Synopsis, &movie.Series, &movie.Path); err != nil {
			fmt.Printf("error scanning movie: %s\n", err)
			return movie, err
		}
	}
	if rows != nil {
		rows.Close()
	}
	fmt.Printf("image: %s, path: %s\n", movie.Image, movie.Path)
	return movie, nil
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

func insertIntoDB(movie Movie) {
	_, err := db.Database.Exec("insert into movies (title, subtitle, director, year, series, length, image, genre, synopsis, path) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", movie.Title, movie.Subtitle, movie.Director, movie.Year, movie.Series, movie.Length, movie.Image, movie.Genre, movie.Synopsis, movie.Path)
	if err != nil {
		fmt.Printf("error adding book: %s\n", err)
	}
}

func validDBEntry(movie Movie) bool {
	rows, err := db.Database.Query("select title from movies where title = ? and subtitle = ? and director = ?", movie.Title, movie.Subtitle, movie.Director)
	if err != nil {
		fmt.Printf("error checking if movie exists: %s\n", err)
		return false
	}
	for rows.Next() {
		fmt.Printf("movie already exists: %s\n", movie.Title)
		return false
	}
	return true
}

func authenticateMovie(w http.ResponseWriter, r *http.Request) (bool, Movie) {
	var movie Movie
	var success bool = false

	movie.Title = r.FormValue("title")
	movie.Subtitle = r.FormValue("subtitle")
	movie.Director = r.FormValue("director")
	if !validDBEntry(movie) {
		return false, movie
	}
	var folderName string = "movies/" + movie.Title + "-" + movie.Subtitle + "," + movie.Director

	movie.Year = r.FormValue("year")
	movie.Series = r.FormValue("series")
	movie.Length = r.FormValue("length")
	movie.Synopsis = r.FormValue("synopsis")
	movie.Genre = r.FormValue("genre")

	imageFile, imageFilename, imageHandler := authenticate.FormVideo("image", r)
	videoFile, videoFilename, videoHandler := authenticate.FormVideo("media", r)

	movie.Path = folderName + "/" + videoHandler.Filename
	movie.Image = folderName + "/" + imageHandler.Filename

	imageFile.Close()
	// videoFile.Close()
	fmt.Printf("%s\n", imageFilename)

	if authenticate.ValidText(movie.Title) &&
		authenticate.ValidText(movie.Subtitle) &&
		authenticate.ValidText(movie.Director) &&
		authenticate.ValidYear(movie.Year) &&
		authenticate.ValidLength(movie.Length) &&
		authenticate.ValidText(movie.Series) &&
		authenticate.ValidText(movie.Synopsis) &&
		authenticate.ValidImage(folderName, imageHandler) &&
		authenticate.ValidVideo(folderName, videoHandler) {
		success = true
	} else {
		return false, movie
	}

	//needs an upload bar to see progress, not sure how to do that
	if !upload.UploadMedia(imageFile, imageFilename, folderName, imageHandler) {
		return false, movie
	}
	if !upload.UploadMedia(videoFile, videoFilename, folderName, videoHandler) {
		return false, movie
	}

	// authenticate.ValidVideo(folderName, videoHandler)

	return success, movie
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

	success, movies := authenticateMovie(w, r)
	if success {
		insertIntoDB(movies)
		fmt.Print("added\n")
	} else {
		fmt.Print("not added\n")
	}

	http.Redirect(w, r, "/addmovie", http.StatusSeeOther)
}
