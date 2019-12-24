// https://www.spoj.com/problems/TEST/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	numbers := []int{}

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		num, _ := strconv.Atoi(text)

		if num == 42 {
			break
		}
		numbers = append(numbers, num)
	}

	for _, v := range numbers {
		fmt.Println(v)
	}
}
