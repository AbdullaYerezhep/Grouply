package handler

import (
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about/" {
		HandleErrors(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		HandleErrors(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("./ui/template/about.html")
	if err != nil {
		HandleErrors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	templ.Execute(w, nil)
}
