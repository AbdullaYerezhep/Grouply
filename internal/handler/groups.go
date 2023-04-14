package handler

import (
	"html/template"
	"net/http"
	"project/internal/data"
)

func Groups(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/groups/" {
		HandleErrors(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		HandleErrors(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	templ, err := template.ParseFiles("./ui/template/groups.html")
	artists := data.GetData()

	templ.Execute(w, artists)
	if err != nil {
		HandleErrors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
