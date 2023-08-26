//A.16.1. Inisialisasi Slice
package main

import "fmt"

func main()  {
	var fruits = []string{"apple","anggur","pisang","melon"}
	fmt.Println(fruits[0])

	// Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya, jika jumlah elemen tidak dituliskan, maka variabel tersebut adalah slice.
	var buahA = []string{"apel","pisang"} // slice
	var buahB = [2]string{"apel","pisang"} // array
	var buahC = [...]string{"apel","pisang"} // array

	fmt.Println(buahA[0])

}
