package content

import (
	"fmt"
	"net/http"
	"text/template"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string

	files = append(files, "../src/Content/templates/navbar.html")
	files = append(files, "../src/Content/templates/footer.html")

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