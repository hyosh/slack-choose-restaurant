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

	//find restaurant with lowest nb_draws
	var draw *Restaurant
	randIndex := rand.Intn(len(r.restaurants))
	maxDraw := 0
	for i := 0; i < len(r.restaurants); i++ {
		restaurant := r.restaurants[i]
		if restaurant.NbDraws > maxDraw {
			maxDraw = restaurant.NbDraws
		}
	}

	if maxDraw == 0 {
		draw = r.restaurants[randIndex]
	} else {
		var candidates []*Restaurant
		//find restaurant with lowest nb_draws
		for i := 0; i < len(r.restaurants); i++ {
			restaurant := r.restaurants[i]
			if restaurant.NbDraws < maxDraw {
				candidates = append(candidates, restaurant)
			}
		}

		if len(candidates) == 0 {
			draw = r.restaurants[randIndex]
		} else {
			randIndex = rand.Intn(len(candidates))
			draw = candidates[randIndex]
		}

	}

	//increment nb_draws
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
