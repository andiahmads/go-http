package goweb

import (
	"net/http"
	"testing"
)

/* serveFile digunakan ketika kita hanya ingin menggunakan static file sesuai dengan yg kita inginkan.
hal ini bisa dilakukan menggunakan function http.ServeFile()
dengan menggunakan function ini, kita bisa menentukan file mana yang ingin kita tulis ke http response.

parameter function http.ServeFile hanya berisi string filename, sehingga tidak bisa menggunakan golang embed.
tetapi jika ingin menggunakan golang embed kita tidak bisa menggunakan http.ServeFile, tetapi harus menggunakan fmt responseWriter,
*/

func ServeFIle(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resource/ok.html")
	} else {
		http.ServeFile(writer, request, "./resource/404.html")

	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFIle),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// //go:embed resource

// var resourceOk string

// //go:embed resource/404.html
// var resourceNotOk string

// func ServeFIleWithEmbed(writer http.ResponseWriter, request *http.Request) {
// 	if request.URL.Query().Get("name") != "" {
// 		fmt.Fprintf(writer, resourceOk)
// 	} else {
// 		fmt.Fprintf(writer, resourceNotOk)

// 	}
// }

// func TestServeFileEmbed(t *testing.T) {
// 	server := http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: http.HandlerFunc(ServeFIleWithEmbed),
// 	}
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		panic(err)
// 	}
// }
