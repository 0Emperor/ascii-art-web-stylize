package backend

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, statusCode int, message string) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Error 500: INTERNAL SERVER ERROR")
		return
	}
	w.WriteHeader(statusCode)
	data := ErrorData{
		StatusCode: statusCode,
		Message:    message,
	}
	tmpl.Execute(w, data)
}
