package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	Port int
	Host string
}

func (s *Server) Start() {
	http.HandleFunc("/restaurants/random", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		responseJson(w, getRandRestaurant())
	})

	http.HandleFunc("/restaurants/reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		restaurants := LoadFromDB()
		restaurants.ReinisialiseDraw()

		responseJson(w, "Restaurants réinitialisés")
	})

	http.HandleFunc("/restaurants/display", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		restaurants := LoadFromDB()

		responseJson(w, fmt.Sprintf("Restaurants : %s", restaurants.DisplayList()))
	})

	http.HandleFunc("/restaurants/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("text")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		deleteRestaurant(name)

		responseJson(w, fmt.Sprintf("Restaurant : %s supprimé", name))
	})

	http.HandleFunc("/restaurants", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		name := r.FormValue("text")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		restaurant := createRestaurant(name)

		responseJson(w, fmt.Sprintf("Restaurant : %s ajouté", restaurant.Name))
	})

	// catch all
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Oops, nothing to see here")
	})

	fmt.Printf("Listening on %s:%d\n", s.Host, s.Port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), nil)
}

func NewServer(port int, host string) *Server {
	return &Server{port, host}
}

func getRandRestaurant() string {
	restaurants := LoadFromDB()
	restaurant := restaurants.Draw()
	return restaurant.Name
}

func createRestaurant(name string) *Restaurant {
	restaurant := CreateRestaurant(name)
	restaurants := LoadFromDB()
	restaurants.Add(restaurant)

	return restaurant
}

func deleteRestaurant(name string) {
	restaurants := LoadFromDB()
	restaurants.Remove(name)
}

type Response struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
}

func responseJson(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{text, "in_channel"}
	json.NewEncoder(w).Encode(response)

	w.WriteHeader(http.StatusOK)
}
