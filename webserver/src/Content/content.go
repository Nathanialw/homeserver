package content

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
	db "webserver/src/DB"

	"github.com/julienschmidt/httprouter"
)

type UserSession struct {
	Movies []db.Movie

	//post history
	//post history json
}

func GenerateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string

	files = append(files, "../src/content/templates/navbar.html")
	files = append(files, "../src/content/templates/footer.html")

	for _, file := range fn {
		files = append(files, fmt.Sprintf("../src/%s/templates/%s.html", file, file))
	}
	for _, file := range files {
		fmt.Print(file)
		fmt.Print("\n")
	}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, fn[0], data)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}

func Movie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// currentPage = r.URL.Path

	fmt.Printf("message received from %s\n"+p.ByName("name"), r.RemoteAddr)
	//strip off the end of the url
	list := strings.TrimPrefix(p.ByName("categoryPath"), "/")
	// fmt.Printf("category: %s\n", list)
	var data UserSession
	// data.LoggedIn = LoginStatus(r)
	// data.Admin = AdminStatus(r)
	// data.Category, _ = GetCategory(list)
	// if data.Category.Category == "" {
	// 	notfound(w, r, p)
	// 	return
	// }

	data.Movies, _ = db.Movies(list)

	GenerateHTML(w, data, "LANMovies")
}
