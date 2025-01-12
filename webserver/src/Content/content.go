package content

import (
	"fmt"
	"net/http"
	"text/template"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, module string, fn ...string) {
	var files []string

	files = append(files, "../templates/General/navbar.html")
	files = append(files, "../templates/General/footer.html")
	files = append(files, "../templates/General/head.html")

	for _, file := range fn {
		files = append(files, fmt.Sprintf("../templates/%s/%s.html", module, file))
	}
	// for _, file := range files {
	// fmt.Print(file)
	// fmt.Print("\n")
	// }
	templates := template.Must(template.ParseFiles(files...))
	_ = templates.ExecuteTemplate(w, fn[0], data)

	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}

//need to do that at the poitn of parsing to get the formatting right
func FormatParagraph(input string) (text string) {

	for i := 0; i < len(input); i++ {
		if input[i] == '.' && i+1 < len(input) && input[i+1] != ' ' && input[i+1] != '.' && input[i+1] != '"' && input[i+1] != '\'' && input[i+2] != '.' {
			input = input[:i+1] + "<br><br>" + input[i+1:]
		}
	}

	text = input

	return text
}
