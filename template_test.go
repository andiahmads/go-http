package goweb

import (
	"html/template"
	"net/http"
	"testing"
)

/*
membuat template pada golang kita perlu memberi tahu nama templatenya.
dan untuk konten yg dinamis kita harus gunakan tanda {{}

*/

func SimpleTempalte(write http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	//cara lebih simple
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(write, "SIMPLE", "hello temple gg wp")
}

func TestTempalte(t *testing.T) {
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(SimpleTempalte),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
