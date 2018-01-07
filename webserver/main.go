package main

import (
	"fmt"
	"io"
	"net/http"
)



func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>17Disney ETL0.2</h1>")
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":5000", nil)
	fmt.Println("17Disney ETL Service")
}