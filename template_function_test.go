package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
didalam template function juga bisa diakses
cara akses function sama seperti mengakses field, namun jika function tersebut memiliki parameter
kita bisa tambahkan parameter ketika memanggil function ditemplatenya dengan syntax functinName()
*/

type Mypage struct {
	Name string
}

func (mypage Mypage) SayHello(name string) string {
	return "hello" + name + ", my name is" + mypage.Name
}

func TemplateFunction(writer http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "andi"}}`))

	t.ExecuteTemplate(writer, "FUNCTION", Mypage{
		Name: "Endi",
	})
}

func TestTempalateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}

func TemplateFunctionGlobal(writer http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", Mypage{
		Name: "Endi",
	})
}

func TestTempalateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
