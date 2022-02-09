package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

/* middleware adalah filter atau juga bisa disebut sebagai interceptor
yg merupakan sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah handler dieksekusi.

namun digolang tidak memiliki fitur middleware, namun karena struktur yang handler yang baik menggunakan interface.
kita bisa membuat middleware sendiri dengan handler.

middleware juga bisa digunakan untuk melakukan error handler.
ketika terjadi panic dihanler, kita bisa melakukan recover dimiddleware, dan mengubah panic tersebut menjadi error response.
dengan ini kita bisa menjaga applikasi kita tidak berhenti berjalan.
*/

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("before execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("after execute Handler")

}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {

		err := recover()
		if err != nil {
			fmt.Println("Terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "error : %s", err)
		}

	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler execute")
		fmt.Fprintf(writer, "hello middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo js ")
		fmt.Fprintf(writer, "hello foo")
		panic("coooooooooooooooooook")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("PANIC")
		panic("coooooooooooooooooook")
	})
	logMiddleware := &LogMiddleware{
		Handler: mux,
	}
	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
