// A.12.2. Operator Perbandingan
package main

import "fmt"

func main()  {
	var value = (((2 + 6) % 3) * 4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("Nilai %d (%t) \n", value, isEqual)
}
