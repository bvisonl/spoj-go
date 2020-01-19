package main

import (
	"encoding/json"
	"fmt"
)

type Case struct {
	Cities []City     `json:"cities"`
	Paths  [][]string `json:"paths"`
}

type City struct {
	Index       int          `json:"index"`
	Name        string       `json:"name"`
	Connections []Connection `json:"connections"`
	Checked     bool         `json:"checked"`
}

type Connection struct {
	CityId int `json:"city"`
	Cost   int `json:"cost"`
}

func main() {

	// Capture cases
	tCases := []Case{}
	captureCases(tCases)

	// Process each case

}

func captureCases(tCases []Case) {

	// Capture # of test cases
	var t int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {

		// Initialize case
		tCase := Case{
			Cities: []City{},
			Paths:  [][]string{},
		}

		// Capture # of cities
		var n int
		fmt.Scanln(&n)

		// Capture cities
		for j := 0; j < n; j++ {

			// Initialize city
			c := City{Index: j + 1, Connections: []Connection{}}
			fmt.Scanln(&c.Name) // Name of the city

			// # of neighbours
			var p int
			fmt.Scanln(&p)

			// Capture connections
			for k := 0; k < p; k++ {
				var nr, cost int
				fmt.Scanln(&nr, &cost)
				c.Connections = append(c.Connections, Connection{CityId: nr, Cost: cost})
			}

			tCase.Cities = append(tCase.Cities, c)
		}

		// # of paths to find
		var r int
		fmt.Scanln(&r)

		for j := 0; j < r; j++ {
			var src, dst string
			fmt.Scanln(&src, &dst)
			tCase.Paths = append(tCase.Paths, []string{src, dst})
		}

		tCases = append(tCases, tCase)

	}

	// Dump the case
	tCasesJson, _ := json.MarshalIndent(tCases, "", "	")
	fmt.Printf("%s\n", tCasesJson)

}
