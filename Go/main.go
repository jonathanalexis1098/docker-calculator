package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Calculate struct {
	Num1 float64 `json:"num1, omitempty"`
	Num2 float64 `json:"num2, omitempty"`
}

type Result struct {
	Value float64
}

func add(w http.ResponseWriter, r *http.Request) {
	var c Calculate
	err := json.NewDecoder(r.Body).Decode(&c)
	log.Println(c)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		var r Result
		r.Value = c.Num1 + c.Num2
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
		log.Println(r)
	}
}

func subtract(w http.ResponseWriter, r *http.Request) {
	var c Calculate
	err := json.NewDecoder(r.Body).Decode(&c)
	log.Println(c)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		var r Result
		r.Value = c.Num1 - c.Num2
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
		log.Println(r)
	}
}

func multiply(w http.ResponseWriter, r *http.Request) {
	var c Calculate
	err := json.NewDecoder(r.Body).Decode(&c)
	log.Println(c)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		var r Result
		r.Value = c.Num1 * c.Num2
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
		log.Println(r)
	}
}

func divide(w http.ResponseWriter, r *http.Request) {
	var c Calculate
	err := json.NewDecoder(r.Body).Decode(&c)
	log.Println(c)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	} else {
		var r Result
		r.Value = c.Num1 / c.Num2
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(r)
		log.Println(r)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/add", add).Methods("POST")
	router.HandleFunc("/subtract", subtract).Methods("POST")
	router.HandleFunc("/multiply", multiply).Methods("POST")
	router.HandleFunc("/divide", divide).Methods("POST")
	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
