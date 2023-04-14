package app

import (
	"fmt"
	"net/http"
	"project/internal/handler"
)

func Run(port string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/groups/", handler.Groups)
	mux.HandleFunc("/about/", handler.About)
	mux.HandleFunc("/search/", handler.SearchArtists)
	mux.HandleFunc("/filter/", handler.FilterArtists)
	styles := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", styles))
	img := http.FileServer(http.Dir("./ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", img))
	fmt.Println("Server is up and listening to port :8080")
	return http.ListenAndServe(port, mux)
}
