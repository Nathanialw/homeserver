package content

import (
	"fmt"
	"net/http"
	"text/template"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, module string, fn ...string) {
	var files []string

	files = append(files, "../content/General/templates/navbar.html")
	files = append(files, "../content/General/templates/footer.html")
	files = append(files, "../content/General/templates/head.html")

	for _, file := range fn {
		files = append(files, fmt.Sprintf("../content/%s/templates/%s.html", module, file))
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
