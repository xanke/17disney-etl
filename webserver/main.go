package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"log"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}


func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>17Disney ETL</h1>")
	reLaunch()
}

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":5000", nil)
	fmt.Println("17Disney ETL Service")
}