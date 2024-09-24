package backend

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		parseINDEX, err := template.ParseFiles("templates/index.html")
		if err != nil {
			ErrorPage(w, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
			return
		}
		parseINDEX.Execute(w, nil)
	case "/assets/":
		if r.URL.Path != "/assets/style.css" {
			ErrorPage(w, http.StatusNotFound, "Page Not Found")
			return
		}
	default:
		ErrorPage(w, http.StatusNotFound, "Page Not Found")
		return
	}
}
