package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

const RESTAURANT_FILE = "restaurants.json"

type Restaurant struct {
	Name    string `json:"name"`
	NbDraws int    `json:"nb_draws"`
}

func CreateRestaurant(name string) *Restaurant {
	return &Restaurant{name, 0}
}

type Restaurants struct {
	restaurants []*Restaurant
}

func (r *Restaurants) Draw() *Restaurant {

	if len(r.restaurants) == 0 {
		return &Restaurant{"Aucun restaurant disponible", 0}
	}

	draw := r.restaurants[rand.Intn(len(r.restaurants))]

	draw.NbDraws++

	//save to json file
	r.EraseDb()

	return draw
}

func (r *Restaurants) EraseDb() {
	//save to json file
	b, err := json.Marshal(r.restaurants)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(RESTAURANT_FILE, b, 0644)
	if err != nil {
		panic(err)
	}

}

func (r *Restaurants) Add(restaurant *Restaurant) {
	exists := false
	for i := 0; i < len(r.restaurants); i++ {
		if r.restaurants[i].Name == restaurant.Name {
			exists = true
			break
		}
	}

	if !exists {
		r.restaurants = append(r.restaurants, restaurant)
	}

	r.EraseDb()
}

func (r *Restaurants) Remove(name string) {
	var newRestaurants []*Restaurant
	for i := 0; i < len(r.restaurants); i++ {
		if r.restaurants[i].Name != name {
			newRestaurants = append(newRestaurants, r.restaurants[i])
		}
	}

	r.restaurants = newRestaurants
	r.EraseDb()
}

func (r *Restaurants) ReinisialiseDraw() {
	for i := 0; i < len(r.restaurants); i++ {
		r.restaurants[i].NbDraws = 0
	}

	r.EraseDb()
}

func (r *Restaurants) DisplayList() string {
	var list string
	for i := 0; i < len(r.restaurants); i++ {
		list += fmt.Sprintf("%s (%d draws) \r", r.restaurants[i].Name, r.restaurants[i].NbDraws)
	}

	return list
}

func LoadFromDB() *Restaurants {

	//load from json file
	b, err := ioutil.ReadFile(RESTAURANT_FILE)
	if err != nil {
		panic(err)
	}

	//unmarshal json
	var restaurants []*Restaurant
	err = json.Unmarshal(b, &restaurants)
	if err != nil {
		panic(err)
	}

	return &Restaurants{restaurants}

}
