// A.16.6. Fungsi append()
/*digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan setelah indeks paling akhir.*/
package main

import "fmt"

func main()  {
	var buah = []string{"pisang","apel","anggur"}
	var cBuah = append(buah,"kates")

	fmt.Println(buah)
	fmt.Println(cBuah)
}
