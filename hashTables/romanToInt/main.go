package main

func romanToInt(s string) int {
	ht := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var sum int

	for i := 0; i < len(s); i++ {
		ch := string(s[i])

		if i+1 < len(s) {
			nextCh := string(s[i+1])

			switch {
			case ch == "I" && nextCh == "V":
				sum += 4
				i++
				continue
			case ch == "I" && nextCh == "X":
				sum += 9
				i++
				continue
			case ch == "X" && nextCh == "L":
				sum += 40
				i++
				continue
			case ch == "X" && nextCh == "C":
				sum += 90
				i++
				continue
			case ch == "C" && nextCh == "D":
				sum += 400
				i++
				continue
			case ch == "C" && nextCh == "M":
				sum += 900
				i++
				continue
			}
		}

		sum += ht[ch]
	}

	return sum
}
