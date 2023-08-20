// A.15.7. Perulangan Elemen Array Menggunakan Keyword for - range
package main

import "fmt"

func main()  {
	var buah = [4]string {
		"apel",
		"nanas",
		"mangga",
		"jeruk",
	}

	for i, buahan := range buah {
		fmt.Printf("elemen %d : %s\n", i, buahan)
	}
}
