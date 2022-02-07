package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/* http merupakan stateles antar client dan server, artinya server tidak akan menyimpan data apapun untuk mengingat setiap requet dari client.
hal ini bertujuan agar mudah melakukan scability disisi server.
lantas bagaimana caranya agar server bisa mengingat client, misalkan case login.
untuk melakukan hal ini kita bisa memanfaatkan cookie.
cookie adalah fitur http (key value)
untuk membuat cookie kita bisa menggunakan function http.setCookie()
*/

func SetCookies(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "x-DOMV"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "succes create cookie")
}

func GetCookies(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("x-DOMV")
	if err != nil {
		fmt.Fprintf(writer, "NO cookie!")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "hello %s", name)
	}
}

func TestRunnigCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookies)
	mux.HandleFunc("/get-cookie", GetCookies)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestSetCookies(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=andi", nil)
	recorder := httptest.NewRecorder()
	SetCookies(recorder, request)

	cookie := recorder.Result().Cookies()
	for _, c := range cookie {
		fmt.Printf("cookie %s:%s \n", c.Name, c.Value)
	}
}

func TestGetCookies(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-DOMV"
	cookie.Value = "andi"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()
	GetCookies(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

}
