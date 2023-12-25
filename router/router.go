package router

import (
	"github.com/aswinayyolath/goapimongoDB/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.ShowAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/movies", controller.DeleteMovies).Methods("DELETE")

	return r
}
