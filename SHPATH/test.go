package main

import (
	"fmt"
)

func main() {

	vals := []int{1, 2, 3}
	testm(vals)
	fmt.Println(vals)
}

func testm(vals []int) {
	vals[0] = 0
}
