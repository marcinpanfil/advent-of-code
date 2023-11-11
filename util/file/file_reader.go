package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return lines
}

func ReadInputAsIntArray() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := []int{}
	for scanner.Scan() {
		intValue, err := strconv.Atoi(scanner.Text())
		if err == nil {
			values = append(values, intValue)
		} else {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return values
}
