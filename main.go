package main

import (
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

func containsDuplicate(nums []int) bool {
	check := map[int]int{}
	for _, val := range nums {
		check[val] += 1
	}
	for _, v := range check {
		if v > 1 {
			return false
		}
	}
	return true
}

func solution(nums []int) int {
	count := map[int]int{}
	for _, val := range nums {
		count[val] += 1
	}

	result := 0
	for key, val := range count {
		if key == val && val > result {
			result = val
		}
	}

	if result == 0 {
		return 0
	}
	return result
}

func main() {

	/*

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


		nums := []int{1, 2, 3, 1}
		fmt.Println(containsDuplicate(nums))
	*/

	//nums := []int{3, 8, 2, 3, 3, 2}
	nums := []int{2, 2, 4, 4, 4}
	fmt.Println(solution(nums))

}
