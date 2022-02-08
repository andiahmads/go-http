package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

func ParseTemplateToFIleHtml(writer http.ResponseWriter, request *http.Request) {
	// t, err := template.ParseFiles("./template/index.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	// t.ExecuteTemplate(writer, "index.gohtml", "Hello goHtml")

	//gunakan go embed
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(writer, "index.gohtml", "Hello goHtml")

}

func TestParseTemplateToFIleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ParseTemplateToFIleHtml(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
