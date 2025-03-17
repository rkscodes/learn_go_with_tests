package iteration

import "strings"

func Repeat(input string, times int) (output string) {
	var sb strings.Builder
	for range times {
		sb.WriteString(input)
	}
	return sb.String()
}

func Repeat2(input string, times int) string {
	var output string
	for i := range times {
		// this proves that strings are stored in read only section of program and is refrenced directly
		// see the allocation results using go tests -bench=. -benchmem
		// you will see
		if i%5 == 0 {
			output += ""
			continue
		}
		output += input
	}
	return output
}
