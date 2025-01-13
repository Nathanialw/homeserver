package content

import (
	"fmt"
	"net/http"
	"reflect"
	"text/template"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, module string, file string, fn ...string) {
	var files []string

	files = append(files, "../templates/General/navbar.html")
	files = append(files, "../templates/General/footer.html")
	files = append(files, "../templates/General/head.html")

	files = append(files, fmt.Sprintf("../templates/%s/%s.html", module, file))

	for _, f := range fn {
		files = append(files, fmt.Sprintf("../templates/General/%s.html", f))
	}

	tmpl := template.New(file + ".html")
	RegisterTemplateFuncs(tmpl)
	tmpl.ParseFiles(files...)
	_ = tmpl.ExecuteTemplate(w, file, data)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

//need to do that at the poitn of parsing to get the formatting right
func FormatParagraph(input string) (text string) {
	for i := 0; i < len(input); i++ {
		if input[i] == '.' && i+1 < len(input) && input[i+1] != ' ' && input[i+1] != '.' && input[i+1] != '"' && input[i+1] != '\'' && input[i+2] != '.' {
			input = input[:i+1] + "<br><br>" + input[i+1:]
		}
		if input[i] == '!' && i+1 < len(input) && input[i+1] != ' ' && input[i+1] != '!' && input[i+1] != '"' && input[i+1] != '\'' && input[i+2] != '!' {
			input = input[:i+1] + "<br><br>" + input[i+1:]
		}
	}

	text = input
	return text
}

func FieldExists(data interface{}, fieldName string) bool {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return false
	}
	field := v.FieldByName(fieldName)
	return field.IsValid()
}

// RegisterTemplateFuncs registers custom template functions
func RegisterTemplateFuncs(t *template.Template) {
	t.Funcs(template.FuncMap{
		"fieldExists": FieldExists,

		//can add more template functions here as needed
	})
}
