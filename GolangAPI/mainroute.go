package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", AnyRoute)
	http.HandleFunc("/getData", getDataParam)
	http.HandleFunc("/getData/1", getDataParam)
	fmt.Println("Server listening on Port 8000")
	http.ListenAndServe(":8000", nil)
}

func AnyRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In anyroute handling /"))
}

func getDataRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In AnyDataRoute handling /"))
}

func getDataWithIdRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In getDataWithIdRoute handling /"))
}

func getDataParam(w http.ResponseWriter, r *http.Request) {
	pattern, _ := regexp.Compile('/getData/(\d+)')
	match:= pattern.FindStringSubmatch(r.URL.Path)

	if len(match) > 0 {
		w.Write([]byte("Value is " + strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]))
	} else {
		w.Write([]byte("No value is present"))
	}
}
