// A.13.4. Pemanfaatan case Untuk Banyak Kondisi
package main

import "fmt"

func main()  {
	var point = 6

	switch point {
	case 8:
		fmt.Println("Saikyooo")

	case 7,6,5,4:
		fmt.Println("Motto ganbare nee~")

	default:
		fmt.Println("Akiramenaide")
	}
}
