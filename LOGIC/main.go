package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Element struct {
	Operand string
	X       int
	Y       int
	Inputs  [][]int
	Outputs [][]int
}

type Circuit struct {
	Elements []Element
}

var n int
var circuits []Circuit

type T struct {
	X int
}

type B struct {
	Ts []T
}

func main() {

	// Header
	fmt.Println("Problem: LOGIC\nLink: https://www.spoj.com/problems/LOGIC/\n\n")

	// Read input
	reader := bufio.NewReader(os.Stdin)

	// Get the number of circuits
	fmt.Println("Enter the number of circuits: ")
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	n, err = strconv.Atoi(line[:len(line)-1])

	if err != nil {
		fmt.Println(err)
		return
	}

	// Capture the circuits
	fmt.Println("Enter the circuit's components and finish with `end`: ")
	for i := 0; i < n; i++ {
		circuit := Circuit{}

	ReadComponent:
		line, err = reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		// Read the elements until the end of the circuit
		for line[:len(line)-1] != "end" {
			// Remove the .. ending
			strings.Replace(line, "..", "", -1)

			// Split the line by space
			components := strings.Split(line[:len(line)-1], " ")

			// Create the element
			element := Element{}

			operandRegex := regexp.MustCompile(`i|&|\||o|!`)
			element.Operand = string(operandRegex.Find([]byte(components[0])))
			element.X, _ = strconv.Atoi(string(components[0][0]))
			element.Y, _ = strconv.Atoi(string(components[0][1]))

			// Capture outputs (-1 since we already have the input)
			for j := 1; j < len(components); j++ {
				x, _ := strconv.Atoi(string(components[j][0]))
				y, _ := strconv.Atoi(string(components[j][1]))
				element.Outputs = append(element.Outputs, []int{x, y})
			}

			circuit.Elements = append(circuit.Elements, element)
			goto ReadComponent
		}

		circuits = append(circuits, circuit)

	}

	fmt.Print(circuits[0].Elements)

}
