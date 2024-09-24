package backend

import (
	"html/template"
	"net/http"

	asciiweb "asciiweb/artistTools"
)

type PageData struct {
	Result string
	Text   string
}

type ErrorData struct {
	StatusCode int
	Message    string
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorPage(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("template")
	if len(text) > 500 {
		ErrorPage(w, http.StatusBadRequest, "Bad Request: Text is too long only less than 500 char accepted")
		return
	}

	if text == "" || banner == "" {
		ErrorPage(w, http.StatusBadRequest, "Bad Request: Text or Banner is emty")
		return
	}

	bannerFile := "banners/" + banner + ".txt"

	if !asciiweb.IsValidASCII(text) {
		ErrorPage(w, http.StatusBadRequest, "Bad Request: Text Contain invalid characters.")
		return
	}

	result := asciiweb.GenerateAsciiArt(text, bannerFile)
	if result == "Error" {
		ErrorPage(w, http.StatusBadRequest, "Bad Request: Banner not found")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	data := PageData{Result: result, Text: text}

	tmpl.Execute(w, data)
}
