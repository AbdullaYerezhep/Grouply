package handler

import (
	// "fmt"
	"html/template"
	"net/http"
	"project/internal/data"
	"project/internal/models"
	"strconv"
	"strings"
)

func SearchArtists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		HandleErrors(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		HandleErrors(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	query := r.FormValue("search-input")
	query = strings.ToLower(query)
	artists := data.GetData()
	var results []models.Artist
	if query != "" {
		for _, artist := range artists {
			if strings.Contains(strings.ToLower(artist.Name), query) {
				results = append(results, artist)
			}

			for _, member := range artist.Members {
				if strings.Contains(strings.ToLower(member), query) {
					results = append(results, artist)
					break
				}
			}

			for location, date := range artist.Event.DatesLocations {
				locationString := location + ":" + date[0]
				if strings.Contains(strings.ToLower(locationString), query) {
					results = append(results, artist)
				}
			}

			if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
				results = append(results, artist)
			}
			if strconv.Itoa(artist.CreationDate) == query {
				results = append(results, artist)
			}
		}
	}
	templ, err := template.ParseFiles("./ui/template/groups.html")
	if err != nil {
		HandleErrors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	templ.Execute(w, results)
}
