// A.15.8. Penggunaan Variabel Underscore _ Dalam for - range
package main

import "fmt"

func main()  {
	var buah = [4]string {
		"apel",
		"nanas",
		"mangga",
		"jeruk",
	}

	for _, buahan := range buah {
		fmt.Printf("nama buah : %s\n",buahan)
	}
}
