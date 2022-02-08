package goweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/* ketika membuat template kita ingin menambahkan banyak data dinamis
hal ini bisa dilakukan dengan cara menggunakan struct atau map.
namum perlu dilakukan perubahan didalam text templatenya.
kita perlu memberi tahu field apa yg kita gunakan untuk mengisi dinamis template.
*/

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/data-template.gohtml"))

	t.ExecuteTemplate(writer, "data-template.gohtml", map[string]interface{}{
		"Title": "template data from golang",
		"Name":  "Andi ahamds",
	})
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()
	TemplateDataMap(recorder, request)
	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
