package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(helloHandler))

	http.ListenAndServe(":8080", mux)

}
func helloHandler(w http.ResponseWriter, req *http.Request) {
	WriteResponse(w, "Hello, world!\n")
}

func WriteResponse(w io.Writer, response string) {
	io.WriteString(w, response)
}
