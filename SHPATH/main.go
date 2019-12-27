// https://www.spoj.com/problems/SHPATH/

package main

import (
	"fmt"
)

type city struct {
	index	   int
	name       string
	neighbours int
	distances  [][]int
	cost       int
	prevCity   *city
}

type test struct {
	n      int // Number of cities
	cities []city
	r      int // Number of paths to find
	paths  [][]string
}

func main() {

	tests := []test{}

	// Number of tests
	var s int
	fmt.Scanln(&s)

	// Capture tests
	for i := 0; i < s; i++ {
		var t test

		fmt.Scanln(&t.n)

		for j := 0; j < t.n; j++ {
			c := city{distances: [][]int{}, cost: 200000, prevCity: nil, index: j+1}

			fmt.Scanln(&c.name)
			fmt.Scanln(&c.neighbours)

			for k := 0; k < c.neighbours; k++ {
				var index, cost int
				fmt.Scanln(&index, &cost)
				c.distances = append(c.distances, []{index, cost})
			}

			t.cities = append(t.cities, c)
		}

		fmt.Scanln(&t.r)

		for j := 0; j < t.r; j++ {
			var src, dst string
			fmt.Scanln(&src, &dst)
			t.paths = append(t.paths, []string{src, dst})
		}

		tests = append(tests, t)
	}

	// Process cases
	for _, t := range tests {

		var cities []city
		copy(cities, t.cities)

		for _, p := range t.paths {

			src := findCity(p[0], cities)
			src.cost = 0
			dst := findCity(p[1], cities)

			dijkstra(src, dst, cities)
		}
	}

}

func dijkstra(src *city, dst *city, cities []city) {
	for len(cities) >= 1 {
		u, cities = cities[len(cities)-1], cities[:len(cities)-1]
		for _, d := range u.distances {
			c := findCityById(d[0], cities)
			alt := u.cost + d[1]
			if(c.cost > alt) {
				c.cost = alt
				c.prevCity = &u
			}
		}
	}
	fmt.Println(dst.cost)
}

func findCityById(index int, cities []city) *city {
	for _, c := range cities {
		if c.index == index {
			return &c
		}
	}
	return nil
}


func findCity(name string, cities []city) *city {
	for _, c := range cities {
		if c.name == name {
			return &c
		}
	}
	return nil
}
