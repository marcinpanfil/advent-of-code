package strings

func ReplaceCharAt(input string, pos int, replacement string) string {
	return input[:pos] + replacement + input[pos+1:]
}
