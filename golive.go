package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	Serve()
}

func Serve() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't get current working directory: ", err)
		return
	}

	http.Handle("/", http.FileServer(http.Dir(wd)))
	log.Printf("We're live at http://0.0.0.0:8108")
	err = http.ListenAndServe("0.0.0.0:8108", nil)
	if err != nil {
		log.Fatal("Can't bind to port 8108:", err)
	}
}
