package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}
func (this *Message) setMessage(message string) {
	this.Message = message
}

var movies = Movies{
	Movie{"Con la muerte en los talones", 2013, "Hitchcok"},
	Movie{"Batman Begins", 1999, "Scorsese"},
	Movie{"A todo gas", 2005, "Juan Antonio"}}

var collection = getSession().DB("curso_go").C("movies")

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:32768")
	if err != nil {
		panic(err)
	}
	return session
}

func responseMovie(w http.ResponseWriter, status int, retults Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(retults)
	return
}

func responseMovies(w http.ResponseWriter, status int, retults []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(retults)
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con Go")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados: ", results)
	}
	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	results := Movie{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, results)
}

func MovieAdd(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var movie_data Movie
	err := decoder.Decode(&movie_data)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	log.Println(movie_data)

	responseMovie(w, 200, movie_data)
}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	new_data := bson.M{"$set": movie_data}
	err = collection.Update(document, new_data)

	if err != nil {
		panic(err)
		w.WriteHeader(404)
		return
	}
	responseMovie(w, 200, movie_data)
}

func MovieDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	result := new(Message)
	result.setStatus("success")
	result.setMessage("La pelicula con ID " + movie_id + " ha sido borrada correctamente.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
