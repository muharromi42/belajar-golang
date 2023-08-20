// A.15.4. Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
package main

import "fmt"

func main()  {
	var angka = [...] int {1,2,3,4,5,10}

	fmt.Println("data array \t:",angka)
	fmt.Println("jumlah elemen \t:",len(angka))
}
