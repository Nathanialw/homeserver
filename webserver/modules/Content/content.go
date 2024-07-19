package content

import (
	"fmt"
	"net/http"
	"text/template"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string

	var footer = "footer"
	var navbar = "navbar"
	files = append(files, fmt.Sprintf("../modules/content/templates/%s.html", footer))
	files = append(files, fmt.Sprintf("../modules/content/templates/%s.html", navbar))
	// tmpl, err := template.ParseFiles("../modules/LANDocs/templates/LANDocs.html")

	// files = append(files, fmt.Sprintf("../%s/templates/%s.html", name, name))
	for _, file := range fn {
		files = append(files, fmt.Sprintf("../modules/%s/templates/%s.html", file, file))
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

func SRCgenerateHTML(w http.ResponseWriter, data interface{}, name string, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("../templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, name, data)
}
