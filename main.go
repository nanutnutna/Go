package main

import (
	"encoding/json"
	"fmt"
)

type movie struct {
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	Rating      float32  `json:"rating"`
	Genres      []string `json:"genres"`
	IsSuperHero bool     `json:"isSuperHero"`
}

func Devisable(start, end int) []int {
	var result = []int{}
	for i := start; i <= end; i++ {
		if i%7 == 0 && i%5 != 0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {

	fmt.Printf("%#v\n", Devisable(2000, 3000))

	body := `[
	{
		"imdbID": "tt4154796",
		"title": "Avengers: Endgame",
		"year": 2019,
		"rating": 8.4,
		"genres": ["Action","Drama"],
		"isSuperHero": true
	},
	{
		"imdbID": "tt4154756",
		"title" : "Avengers: Infinity War",
		"year" : 2018,
		"rating": 8.4,
		"genres": ["Action", "Sci-Fi"],
		"isSuperHero": true
	}
	]`

	ms := []movie{}
	err := json.Unmarshal([]byte(body), &ms)
	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v\n", ms)

	for _, v := range ms {
		fmt.Println(v.Title)
	}

}
