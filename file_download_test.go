package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

/*
pada golang untuk mendownload file kita bisa menggunakan FileServer dan ServeFILE
Dan jika ingin memaksa mendownload (tanpa dirender oleh browser) kita bisa menggunakan header Content-Disposition
*/

func DownloadFile(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")
	if fileName == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "BAD Request")
		return
	}
	writer.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(writer, request, "./resource/"+fileName)
}

func TestDownloadFile(t *testing.T) {
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
