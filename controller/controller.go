package controller

import (
	"encoding/json"
	"net/http"

	"github.com/aswinayyolath/goapimongoDB/dbhelper"
	"github.com/aswinayyolath/goapimongoDB/model"
	"github.com/gorilla/mux"
)

func ShowAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := dbhelper.GetAllMovie()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	dbhelper.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	dbhelper.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	dbhelper.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	dbhelper.DeleteAllMovie()
	json.NewEncoder(w).Encode("Deleted all movies")
}
