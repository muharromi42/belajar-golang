// A.13.5. Kurung Kurawal Pada Keyword case & default
package main

import "fmt"

func main() {
	var point = 61

	switch point {
	case 8:
		fmt.Println("Saikyoo")

	case 7,6,5,4:
		fmt.Println("Motto Ganbare Nee~")

	default:
		{
			fmt.Println("Akiramenaide")
			fmt.Println("sugi wa seikou suru deshou")
		}
	}
}
