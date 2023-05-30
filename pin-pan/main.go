package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var result string

		if i % 3 == 0 {
			result += "Pin "
		}
		if i % 5 == 0 {
			result += "Pan"
		}

		if len(result) > 0 {
			fmt.Println(result)
		} else {
			fmt.Println(i)
		}
	}
}