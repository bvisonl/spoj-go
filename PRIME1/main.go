// https://www.spoj.com/problems/PRIME1/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type tCase struct {
	m int
	n int
}

func main() {

	// T cases
	reader := bufio.NewReader(os.Stdin)
	tStr, _ := reader.ReadString('\n')
	tStr = strings.TrimSuffix(tStr, "\n")
	t, _ := strconv.Atoi(tStr)

	tCases := []tCase{}

	// Capture cases
	for i := 1; i <= t; i++ {
		casesStr, _ := reader.ReadString('\n')
		casesStr = strings.TrimSuffix(casesStr, "\n")

		cases := strings.Split(casesStr, " ")

		m, _ := strconv.Atoi(cases[0])
		n, _ := strconv.Atoi(cases[1])

		tCases = append(tCases, tCase{m: m, n: n})

	}

	// Process cases
	for _, c := range tCases {
		for i := c.m; i <= c.n; i++ {
			if isPrime(i) == true {
				fmt.Println(i)
			}
		}
		if tCases[len(tCases)-1] != c {
			fmt.Println("")
		}
	}

}

func isPrime(num int) bool {

	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	// Trial division
	r := int(math.Sqrt(float64(num)))

	for i := 3; i <= r; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true

}
