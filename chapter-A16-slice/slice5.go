// A.16.5. Fungsi cap()
/*Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice. Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len, tapi bisa berubah seiring operasi slice yang dilakukan.*/
package main

import "fmt"

func main()  {
	var buah = []string{"apel","nanas","anggur","pisang"}
	fmt.Println(len(buah)) //4
	fmt.Println(cap(buah)) //4

	var aBuah = buah[0:3]
	fmt.Println(len(aBuah)) //3
	fmt.Println(cap(aBuah)) //4

	var bBuah = buah[1:4]
	fmt.Println(len(bBuah)) //3
	fmt.Println(cap(bBuah)) //3

}
