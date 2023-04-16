package handler

import (
	"html/template"
	"net/http"
)

func HandleErrors(w http.ResponseWriter, message string, code int) {
	templ, err := template.ParseFiles("./ui/template/error.html")
	if err != nil {
		http.Error(w, "Invalid template", http.StatusInternalServerError)
		return
	}
	errorP := struct {
		Code    int
		Messege string
	}{
		Code:    code,
		Messege: message,
	}
	w.WriteHeader(code)
	templ.Execute(w, errorP)
}
