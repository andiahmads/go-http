package goweb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	msg := request.URL.Query().Get("msg")
	fmt.Fprintf(w, " hello %s %s\n", name, msg)

}

func TestQueryParamter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=andi&msg=ganteng", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)
	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	bodyToString := string(body)
	fmt.Println(bodyToString)
}

func MultipleParamterWithMap(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["names"]
	fmt.Fprintf(writer, strings.Join(names, ","))
}

func TestMultipleQueryParamterWithMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?names=andi&names=joni", nil)
	recorder := httptest.NewRecorder()

	MultipleParamterWithMap(recorder, request)
	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	bodyToString := string(body)
	fmt.Println(bodyToString)
}
