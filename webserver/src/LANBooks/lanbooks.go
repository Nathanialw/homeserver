package lanbooks

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

type PageData struct {
	Title string
	Body  string
}

type Book struct {
	Uid      string
	Title    string
	Subtitle string
	Author   string
	Image    string
	Path     string
}

type BookList struct {
	User     user.Session
	NotEmpty bool
	Books    []Book
	Back     string
	Add      string
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	createDB()

	var err error
	var data BookList

	data.Books, err = getAll()
	if len(data.Books) > 0 {
		data.NotEmpty = true
		fmt.Printf("The director is %s.\n", data.Books[0].Author)
	} else {
		data.NotEmpty = false
		fmt.Printf("none found\n")
	}

	data.Back = "/"
	data.Add = "/addbook"

	content.GenerateHTML(w, data, "LANBooks", "home")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getAll() (threads []Book, err error) {
	rows, err := db.Database.Query("select title, author, image from books")
	for rows.Next() {
		th := Book{}
		if err = rows.Scan(&th.Title, &th.Author, &th.Image); err != nil {
			fmt.Printf("%s", err)
			return
		}
		threads = append(threads, th)
		fmt.Printf("name: %s\n", th.Title)
	}
	if rows != nil {
		rows.Close()
	}
	fmt.Println("retrieving all")
	return
}

func createDB() {
	stmt, err := db.Database.Prepare("CREATE TABLE IF NOT EXISTS books (uid INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, subtitle TEXT, author TEXT, image TEXT, path TEXT)")
	if err != nil {
		log.Fatalf("failed to prepare statement: %v\n", err)
	}
	_, execErr := stmt.Exec()
	if execErr != nil {
		log.Fatalf("Failed to execute table creation statement: %v", execErr)
	}
}

func RetrieveFromDB(title string) (Book, error) {
	var movie Book
	rows, err := db.Database.Query("select * from books where title = ?", title)
	if err != nil {
		fmt.Printf("error retrieving movie: %s\n", err)
		return movie, err
	}
	for rows.Next() {
		if err = rows.Scan(&movie.Uid, &movie.Title, &movie.Subtitle, &movie.Author, &movie.Image, &movie.Path); err != nil {
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

func Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var err error
	var data Book

	//need to mkae sure th "movieID" actually exists so it can 404 if it doesn't
	data, err = RetrieveFromDB(p.ByName("bookID"))

	fmt.Printf("path: %s\n", data.Path)
	http.Redirect(w, r, data.Path, http.StatusSeeOther) //open the pdf directly

	// content.GenerateHTML(w, data, "LANBooks", "book") //open the pdf as an element in a page

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)

	var data BookList
	// user.Session.LoggedIn = LoginStatus(r)
	// user.Session.Admin = AdminStatus(r)

	data.Back = "/books"
	data.Add = ""

	content.GenerateHTML(w, data, "LANBooks", "add")
}

func insertIntoDB(movie Book) {
	_, err := db.Database.Exec("insert into books (title, subtitle, author, image, path) values (?, ?, ?, ?, ?)", movie.Title, movie.Subtitle, movie.Author, movie.Image, movie.Path)
	if err != nil {
		fmt.Printf("error adding book: %s\n", err)
	}
}

func validDBEntry(movie Book) bool {
	rows, err := db.Database.Query("select title from books where title = ? and subtitle = ? and author = ?", movie.Title, movie.Subtitle, movie.Author)
	if err != nil {
		fmt.Printf("error checking if Book exists: %s\n", err)
		return false
	}
	for rows.Next() {
		fmt.Printf("Book already exists: %s\n", movie.Title)
		return false
	}
	return true
}

func Authenticate(w http.ResponseWriter, r *http.Request) (bool, Book) {
	var book Book
	var success bool = false

	book.Title = r.FormValue("title")
	book.Subtitle = r.FormValue("subtitle")
	book.Author = r.FormValue("author")
	if !validDBEntry(book) {
		return false, book
	}
	var folderName string = "/mnt/media/books/" + book.Title + "_" + book.Subtitle + "_" + book.Author

	fmt.Printf("folder name: %s\n", folderName)
	// imageFile, imageHandler := authenticate.FormMedia("image", r)
	videoFile, videoHandler := authenticate.FormMedia("media", r)

	// movie.Image = folderName + "/" + imageHandler.Filename
	book.Path = folderName + "/" + videoHandler.Filename

	if authenticate.ValidText(book.Title) &&
		authenticate.ValidText(book.Subtitle) {
		// authenticate.ValidImage(movie.Image, imageHandler) &&

		success = true
	} else {
		return false, book
	}

	//needs an upload bar to see progress, not sure how to do that
	// if !upload.UploadMedia(imageFile, folderName, imageHandler) {
	// 	return false, movie
	// }
	if !upload.UploadMedia(videoFile, folderName, videoHandler) {
		return false, book
	}

	// authenticate.ValidVideo(folderName, videoHandler)

	// imageFile.Close()
	videoFile.Close()
	return success, book
}

func SubmitBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	success, books := Authenticate(w, r)
	if success {
		insertIntoDB(books)
		fmt.Print("added\n")
	} else {
		fmt.Print("not added\n")
	}

	http.Redirect(w, r, "/books", http.StatusSeeOther)
}
