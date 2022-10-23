package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type inData []struct {
	RecipeName  string `json:"recipe_name"`
	Serves      string `json:"serves"`
	CookingTime string `json:"cooking_time"`
	Difficulty  string `json:"difficulty"`
	Ingredients []struct {
		Quantity string `json:"quantity"`
		Name     string `json:"name"`
	} `json:"ingredients"`
	Directions  string   `json:"directions"`
	Preferences []string `json:"preferences"`
	Image       string   `json:"image"`
	RecipeUrls  string   `json:"recipe_urls"`
}

type outData []struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Kcal        int    `json:"kcal"`
	Image       string `json:"image"`
	URL         string `json:"url"`
	Persons     int    `json:"persons"`
	CookingTime string `json:"cookingTime"`
	Ingredients []struct {
		Ingredient struct {
			Name  string `json:"name"`
			Stock int    `json:"stock"`
		} `json:"ingredient,omitempty"`
		Quantity string `json:"quantity"`
	} `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

var (
	id inData
	od outData
)

func main() {
	file, _ := os.Open("BellRezepte_full.json")
	bv, _ := ioutil.ReadAll(file)

	var rList outData = outData{}

	defer file.Close()

	json.Unmarshal(bv, &id)

	fmt.Printf("%+v", id)

	for i, r := range id {

		person, _ := strconv.ParseInt(r.Serves, 2, 6)

		var igl []struct {
			Ingredient struct {
				Name  string "json:\"name\""
				Stock int    "json:\"stock\""
			} "json:\"ingredient,omitempty\""
			Quantity string "json:\"quantity\""
		}

		for i, ig := range r.Ingredients {

			fmt.Println(i)
			fmt.Println(ig)

			igl = append(igl, struct {
				Ingredient struct {
					Name  string "json:\"name\""
					Stock int    "json:\"stock\""
				} "json:\"ingredient,omitempty\""
				Quantity string "json:\"quantity\""
			}{
				Ingredient: struct {
					Name  string "json:\"name\""
					Stock int    "json:\"stock\""
				}{
					Name:  ig.Name,
					Stock: 0,
				},
				Quantity: ig.Quantity,
			})
		}

		rList = append(rList, struct {
			ID          int    "json:\"id\""
			Name        string "json:\"name\""
			Category    string "json:\"category\""
			Kcal        int    "json:\"kcal\""
			Image       string "json:\"image\""
			URL         string "json:\"url\""
			Persons     int    "json:\"persons\""
			CookingTime string "json:\"cookingTime\""
			Ingredients []struct {
				Ingredient struct {
					Name  string "json:\"name\""
					Stock int    "json:\"stock\""
				} "json:\"ingredient,omitempty\""
				Quantity string "json:\"quantity\""
			} "json:\"ingredients\""
			Instructions []string "json:\"instructions\""
		}{
			ID:          i,
			Name:        r.RecipeName,
			Category:    "General",
			Kcal:        500,
			Image:       "https://www.bell.ch" + r.Image,
			URL:         r.RecipeUrls,
			Persons:     int(person),
			CookingTime: r.CookingTime,
			Instructions: []string{
				r.Directions,
			},
			Ingredients: igl,
		})

	}

	u, _ := json.Marshal(rList)

	_ = ioutil.WriteFile("out.json", u, 0640)

	// _ = ioutil.WriteFile("test.json", file, 0644)

}
