package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DEF_FILE_NAME = "input.txt"

func ReadInput() []string {
	file, err := os.Open(DEF_FILE_NAME)
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
	file, err := os.Open(DEF_FILE_NAME)
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

func ReadInputAsSingle() string {
	file, err := os.ReadFile(DEF_FILE_NAME)
	if err != nil {
		fmt.Println(err)
	}
	return strings.ReplaceAll(string(file), "\r", "")
}
