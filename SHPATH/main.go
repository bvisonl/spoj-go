// https://www.spoj.com/problems/SHPATH/

package main

import (
	"fmt"
)

type distance struct {
	index int
	cost int
}

type city struct {
	name string
	distances []distance
}

type test struct {
	n int // Number of cities
	cities []city
	r int // Number of paths to find
	paths [][]string

}

func main()  {

	var tests = []test

	// Number of tests
	var s int
	fmt.Scanln(&s)

	// Capture tests
	for i := 0; i < s; i++ {
		var n int
		fmt.Scanln(&n)

		var name string
		fmt.Scanln(&name)

		var p int
		fmt.Scanln(&p)

		for j := 0; j < p; j++ {
			var nr int
			var cost int
			fmt.Scanln(&nr, &cost)

		}



	}

	var n int

	fmt.Println(m, n)

	a := distance{index: 1, cost: 2}
	b := []distance{}
	b = append(b, a)
	c := city{name: "Hi", distances: b}

	fmt.Println(c)

}
