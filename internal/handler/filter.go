package handler

import (
	"html/template"
	"net/http"
	"project/internal/data"
	"project/internal/models"
	"strconv"
	// "project/internal/models"
	// "strconv"
)

func FilterArtists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter/" {
		HandleErrors(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		HandleErrors(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	templ, err := template.ParseFiles("./ui/template/groups.html")
	if err != nil {
		HandleErrors(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	creationAfter, _ := strconv.Atoi(r.FormValue("creation_after"))
	creationBefore, _ := strconv.Atoi(r.FormValue("creation_before"))

	firstAlbumAfter, _ := strconv.Atoi(r.FormValue("first_album_after"))
	firstAlbumBefore, _ := strconv.Atoi(r.FormValue("first_album_before"))

	location := r.FormValue("location")

	numberOfMembersString := r.Form["member"]
	var numberOfMembers []int

	for _, number := range numberOfMembersString {
		memberNumber, _ := strconv.Atoi(number)
		numberOfMembers = append(numberOfMembers, memberNumber)
	}

	var filteredArtists []models.Artist
	artists := data.GetData()

	for _, artist := range artists {
		artistAlbumDate, _ := strconv.Atoi(artist.FirstAlbum[6:])
		if location == "" && numberOfMembers == nil {
			for i := creationAfter; i <= creationBefore; i++ {
				for j := firstAlbumAfter; j <= firstAlbumBefore; j++ {
					if (artist.CreationDate == i) && (artistAlbumDate == j) {
						filteredArtists = append(filteredArtists, artist)
					}
				}
			}
		} else if location != "" && numberOfMembers == nil {
			for i := creationAfter; i <= creationBefore; i++ {
				for j := firstAlbumAfter; j <= firstAlbumBefore; j++ {
					for key := range artist.Event.DatesLocations {
						if (artist.CreationDate == i) && (artistAlbumDate == j) && (location == key) {
							filteredArtists = append(filteredArtists, artist)
						}
					}
				}
			}
		} else if location == "" && numberOfMembers != nil {
			for i := creationAfter; i <= creationBefore; i++ {
				for j := firstAlbumAfter; j <= firstAlbumBefore; j++ {
					for k := 0; k <= len(numberOfMembers)-1; k++ {
						if (artist.CreationDate == i) && (artistAlbumDate == j) && (len(artist.Members) == numberOfMembers[k]) {
							filteredArtists = append(filteredArtists, artist)
						}
					}
				}
			}
		} else {
			for i := creationAfter; i <= creationBefore; i++ {
				for j := firstAlbumAfter; j <= firstAlbumBefore; j++ {
					for k := 0; k <= len(numberOfMembers)-1; k++ {
						for key := range artist.Event.DatesLocations {
							if (artist.CreationDate == i) && (artistAlbumDate == j) && (location == key) && (len(artist.Members) == numberOfMembers[k]) {
								filteredArtists = append(filteredArtists, artist)
							}
						}
					}
				}
			}
		}
	}

	templ.Execute(w, filteredArtists)
}
