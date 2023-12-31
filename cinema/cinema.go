package cinema

import (
	"fmt"
	"strings"
)

type voter interface {
	addVote(rating float64)
}

func vote(v voter, rating float64) {
	v.addVote(rating)
}

type movie struct {
	name        string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m *movie) addVote(rating float64) {
	m.votes = append(m.votes, rating)
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
	fmt.Printf("\t\t\t     %s\n", m.genres[1])
}

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var total int32
	for _, value := range ar {
		total += value
	}
	return total

}

func compareTriplets(a []int32, b []int32) []int32 {
	var alice, bob int32
	var result []int32
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] > b[i] {
			alice += 1
		} else if a[i] < b[i] {
			bob += 1
		}
	}
	result = append(result, alice, bob)
	return result
}

func WordCount(s string) map[string]int {
	r := map[string]int{}
	for _, v := range strings.Fields(s) {
		r[v] += 1
	}
	return r
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

	var price int = 9999
	var addr *int = &price
	fmt.Println(price, addr)
	*addr = 9400
	fmt.Println(price, addr)
	v := *addr
	fmt.Println(v)

	eg := &movie{
		name:        "Avegers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	//eg.addVote(8)
	vote(eg, 8)
	fmt.Println("votes:", eg.votes)

	s := "If it look like a duck swims like a duck and quacks like a duck then it probably is a duck"
	w := WordCount(s)
	fmt.Printf("%#v\n", w)
}
