package main

import (
	"fmt"
	"math"
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

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

func (c cube) Volume() float64 {
	return c.edge * c.edge * c.edge
}

func (t triangularPrism) Volume() float64 {
	return 0.5 * t.base * t.attitude * t.height
}

func maxArea(height []int) int {
	result := -1.0
	for i := 0; i < len(height); i++ {
		w := 0
		for j := i + 1; j < len(height); j++ {
			w++
			min := math.Min(float64(height[i]), float64(height[j]))
			area := min * float64(w)
			//fmt.Printf("x %#v, y %#v , wid %#v, area %#v\n", height[i], height[j], w, area)
			if area > result {
				result = area
			}
		}
	}
	return int(result)
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

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))

}
