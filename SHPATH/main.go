// https://www.spoj.com/problems/SHPATH/

package main

import (
	"fmt"
)

type city struct {
	name       string
	neighbours int
	distances  map[int]int
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
			c := city{distances: map[int]int{}}

			fmt.Scanln(&c.name)
			fmt.Scanln(&c.neighbours)

			for k := 0; k < c.neighbours; k++ {
				var index, cost int
				fmt.Scanln(&index, &cost)
				c.distances[index] = cost
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

	fmt.Println(tests)

}
