package data

import (
	"encoding/json"
	"io"
	"net/http"
	"project/internal/models"
)

const (
	urlArtists   = "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetData() []models.Artist {
	artistReq, errA := http.Get(urlArtists)
	if errA != nil || artistReq.Status == http.StatusText(http.StatusOK) {
		panic(errA)
	}
	relationReq, errR := http.Get(urlRelations)
	if errR != nil || relationReq.Status == http.StatusText(http.StatusOK) {
		panic(errR)
	}

	data1, err := io.ReadAll(artistReq.Body)
	if err != nil {
		panic(err)
	}
	data2, err := io.ReadAll(relationReq.Body)
	if err != nil {
		panic(err)
	}

	var artists []models.Artist
	var d models.Relations

	json.Unmarshal(data1, &artists)
	json.Unmarshal(data2, &d)
	relations := d.Index

	for i := 0; i < len(artists); i++ {
		artists[i].Event = relations[i]
	}

	return artists

	// return artists, relations
}
