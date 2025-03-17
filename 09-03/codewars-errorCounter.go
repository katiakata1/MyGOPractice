package main

import "fmt"

func printerError(s string) string {
	var errorCount int

	for _, char := range s {
		if char < 'a' || char > 'm' {
			errorCount++
		}
	}

	return fmt.Sprintf("%d/%d", errorCount, len(s))
}

func main() {
	fmt.Println(printerError("aaabbbbhaijjjm"))
	fmt.Println(printerError("aaaxbbbbyyhwawiwjjjwwm"))
}

// a and m have their own codes in Unicode meaning that go is aware of range
