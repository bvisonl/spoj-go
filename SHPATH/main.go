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

func main()  {
	var m int
	var n int

	fmt.Scanln(&m, &n)
	fmt.Println(m, n)

	a := distance{index: 1, cost: 2}
	b := []distance{}
	b = append(b, a)
	c := city{name: "Hi", distances: b}

	fmt.Println(c)

}
