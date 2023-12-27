package main

import "fmt"

type movie struct {
	name        string
	year        int
	rating      float32
	genres      []string
	isSuperHero bool
}

func emote(rating float64) string {
	switch {
	case rating < 5.0:
		return "Dis"
	case rating >= 5.0 && rating < 7.0:
		return "Nor"
	case rating >= 7.0 && rating < 10.0:
		return "Good"
	default:
		return "Nil"
	}
}

func (m movie) info() {
	fmt.Printf("%v", m.name)
	fmt.Printf("(%d) ", m.year)
	fmt.Printf("- (%.2f)\n", m.rating)
	fmt.Printf("Genres:\n")
	fmt.Printf("\t\t\t     %s\n", m.genres[0])
	fmt.Printf("\t\t\t     %s", m.genres[1])
}

func main() {

	var mvs []movie
	movie1 := movie{name: "Aveg: EN", year: 2019, rating: 8.4, genres: []string{"Ac", "Dra"}, isSuperHero: true}
	movie2 := movie{name: "Aveg: In", year: 2018, rating: 8.4, genres: []string{"Ac", "Sc"}, isSuperHero: true}
	mvs = append(mvs, movie1, movie2)
	for _, v := range mvs {
		fmt.Printf("%#v\n", v)
	}

	ae := movie{
		name:        "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	ae.info()
}
