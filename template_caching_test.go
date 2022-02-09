package goweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
idealnya ketika kita menggunakan go template kita harus melakukan parsing hanya 1 kali saja.
dan ini bisa dilakukan dengan teknik chacing(disimpan dimemory),  sehingga kita tidak perlu melakukan parsing lagi.
hal ini mempengaruhi kinerja dari web.
*/

//go:embed templates/*.gohtml
var templatesa embed.FS

var myTemplates = template.Must(template.ParseFS(templatesa, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Caching")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
