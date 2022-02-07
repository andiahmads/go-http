package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")
	writer.Header().Add("X-Powered-By", "andiahmads")
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()
	RequestHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	bodyToStrinf := string(body)
	fmt.Println(bodyToStrinf)

}

func GetHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "andiahmads")
	fmt.Fprint(writer, "ok")
}

func TestGetHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()
	GetHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	bodyToStrinf := string(body)
	fmt.Println(bodyToStrinf)
	fmt.Println(response.Header.Get("X-Powered-By"))

}
