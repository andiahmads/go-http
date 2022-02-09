package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/* untuk melakukan import, kita bisa gunakan perintah berikut:
{{template "name"}}, artinya kita akan mengimport template "name" tanpa memberikan data apapun.
{{template "name".Value}},artinya kita akan mengimport template "name" dengan memberikan data value.
*/

func TemplateLayout(write http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	t.ExecuteTemplate(write, "layout", map[string]interface{}{
		"Name":  "eko",
		"Title": "Template layout",
	})

}

func TestTempalateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
