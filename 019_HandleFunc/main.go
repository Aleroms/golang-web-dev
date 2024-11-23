package main

import (
	"fmt"
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		// io.WriteString(res, "this is the home route")
		fmt.Fprintln(res,"This is also available to use")
	})

	http.ListenAndServe(":8080", nil)
}
