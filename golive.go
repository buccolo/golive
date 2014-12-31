package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't get current working directory: ", err)
	}

	http.Handle("/", http.FileServer(http.Dir(wd)))
	log.Printf("We're live at http://0.0.0.0:8108")
	err = http.ListenAndServe(":8108", nil)
	if err != nil {
		log.Fatal("Can't bind to port 8108:", err)
	}
}
