package main

import (
	"fmt"
	"log"
	"strconv"
)

type Element struct {
	Operand  string
	Outputs  []string
	Computed int
	Data     []int
}

func (e *Element) compute() int {
	if e.Computed != -1 {
		return e.Computed
	}

	if e.Operand == "&" || e.Operand == "|" {
		if len(e.Data) != 2 {
			return -1
		} else {
			if e.Operand == "&" {
				e.Computed = e.Data[0] & e.Data[1]
			} else {
				e.Computed = e.Data[0] | e.Data[1]
			}
		}
	}

	if e.Operand == "i" || e.Operand == "o" {
		return -1
	}

	if e.Operand == "!" {
		e.Computed = ^e.Data[0]
	}

	return e.Computed
}

type Circuit struct {
	Elements [10][10]Element
	Inputs   []string
	Outputs  []string
}

func (c *Circuit) process(data string) string {

	// Break the data
	dataPos := []rune(data)

	if len(dataPos) != len(c.Inputs) {
		log.Fatalln("The data and the inputs don't match")
	}

	var newInputs []string

	for i := 0; i < len(c.Inputs); i++ {

		// Sit in the input position
		inputPos := []rune(c.Inputs[i])

		x, _ := strconv.Atoi(string(inputPos[0]))
		y, _ := strconv.Atoi(string(inputPos[1]))

		// Look for the element that this input should feed
		el := c.Elements[x][y]
		el.Data = append(el.Data, int(dataPos[i]))
		el.compute()

		newInputs = append(newInputs, string(x)+string(y))
	}

	fmt.Printf("%+v\n", c.Elements[0][1])

	return ""

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

	// // Read input
	// reader := bufio.NewReader(os.Stdin)

	// // Get the number of circuits
	// fmt.Println("Enter the number of circuits: ")
	// line, err := reader.ReadString('\n')

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// n, err = strconv.Atoi(line[:len(line)-1])

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Capture the circuits
	// fmt.Println("Enter the circuit's components and finish with `end`: ")
	// for i := 0; i < n; i++ {
	// 	circuit := Circuit{}

	// ReadComponent:
	// 	line, err = reader.ReadString('\n')

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	// Read the elements until the end of the circuit
	// 	for line[:len(line)-1] != "end" {
	// 		// Remove the .. ending
	// 		strings.Replace(line, "..", "", -1)

	// 		// Split the line by space
	// 		components := strings.Split(line[:len(line)-1], " ")

	// 		// Create the element
	// 		x, _ := strconv.Atoi(string(components[0][0]))
	// 		y, _ := strconv.Atoi(string(components[0][1]))
	// 		circuit.Elements[x][y] = Element{Computed: -1}

	// 		operandRegex := regexp.MustCompile(`i|&|\||o|!`)
	// 		circuit.Elements[x][y].Operand = string(operandRegex.Find([]byte(components[0])))

	// 		if circuit.Elements[x][y].Operand == "i" {
	// 			circuit.Inputs = append(circuit.Inputs, string(x)+string(y))
	// 		}

	// 		if circuit.Elements[x][y].Operand == "o" {
	// 			circuit.Outputs = append(circuit.Outputs, string(x)+string(y))
	// 		}

	// 		// Capture outputs (-1 since we already have the input)
	// 		for j := 1; j < len(components); j++ {
	// 			x, _ := strconv.Atoi(string(components[j][0]))
	// 			y, _ := strconv.Atoi(string(components[j][1]))
	// 			circuit.Elements[x][y].Outputs = append(circuit.Elements[x][y].Outputs, string(x)+string(y))
	// 		}

	// 		goto ReadComponent
	// 	}

	// 	circuits = append(circuits, circuit)

	// }

	circuit := Circuit{}
	in1 := Element{Operand: "i", Outputs: []string{"01"}}
	circuit.Elements[0][0] = in1
	circuit.Inputs = append(circuit.Inputs, "00")
	in2 := Element{Operand: "i", Outputs: []string{"01"}}
	circuit.Elements[1][0] = in2
	circuit.Inputs = append(circuit.Inputs, "10")
	el1 := Element{Operand: "&", Outputs: []string{"11"}}
	circuit.Elements[0][1] = el1
	o1 := Element{Operand: "o"}
	circuit.Outputs = append(circuit.Outputs, "01")
	circuit.Elements[1][1] = o1
	circuits = append(circuits, circuit)

	// fmt.Println(circuits[0].Elements)

	fmt.Println(circuits[0].process("11"))

	// for _, c := range circuits {
	// 	outs := c.process("00")
	// }

}
