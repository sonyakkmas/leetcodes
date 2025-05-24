package inttoroman

import "fmt"

func main() {
	num := 3749

	romanToInt := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string

	for _, entry := range romanToInt {
		for num >= entry.value {
			result += entry.symbol
			num -= entry.value
		}
	}

	fmt.Println(result)
}
