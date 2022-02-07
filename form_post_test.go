package goweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/* semua data formpost yang dikirim dari client, secara otomatis akan disimpan dalam attribut Request.PostForm dengan menggunakan method Request.ParseForm()
method ini digunakan untuk melakukan parsing data body apakah bisa diparsing menjadi form atau tidak.
*/

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	firstname := request.PostForm.Get("firstname")
	lastname := request.PostForm.Get("lastname")
	fmt.Fprintf(writer, "%s %s", firstname, lastname)
}

func TestPostForm(t *testing.T) {
	requsetBody := strings.NewReader("firstname=andi&lastname=ahmad")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requsetBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()
	FormPost(recorder, request)

	response := recorder.Result()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
