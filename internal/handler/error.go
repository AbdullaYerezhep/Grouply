package handler

import (
	"net/http"
)

func HandleErrors(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
}
