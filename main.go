package main

import (
	"fmt"
	"net/http"
	"os"

	backend "asciiweb/Backend"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Server run by default in 8080 port, u don't have to specify any thing\n\nEx: go run .")
		return
	}
	http.HandleFunc("/", backend.MainPage)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/ascii-art", backend.AsciiArt)
	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Failed")
		return
	}
}
