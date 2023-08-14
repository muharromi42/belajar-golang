// A.13.1. Seleksi Kondisi Menggunakan Keyword if, else if, & else
package main

import "fmt"

func main()  {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. Nilai anda : %d\n", point)
	}
}
