package handler

import (
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./ui/template/about.html")
	if err != nil {
		HandleErrors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)
}
