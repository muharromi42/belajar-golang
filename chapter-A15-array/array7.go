// A.15.6. Perulangan Elemen Array Menggunakan Keyword for
package main

import "fmt"

func main()  {
	var buah = [4]string {
		"apel",
		"anggur",
		"nanas",
		"jeruk",
	}

	for i := 0; i < len(buah); i++ {
		fmt.Printf("elemen %d : %s\n",i,buah[i])
	}
}
