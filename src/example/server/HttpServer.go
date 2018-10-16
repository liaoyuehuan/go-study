package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Write([]byte("<h1>hello world</h1>"))
		fmt.Fprintf(writer, "<h1>%s : Hello world</h1>", request.FormValue("name"))
	})

	http.ListenAndServe("0.0.0.0:8081", nil)
}
