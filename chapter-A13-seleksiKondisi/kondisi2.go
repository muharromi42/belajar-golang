// A.13.2. Variabel Temporary Pada if - else
/*Variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi di mana ia ditempatkan saja. Penggunaan variabel ini membawa beberapa manfaat, antara lain:

    Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
    Kode menjadi lebih rapi
    Ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan di dalam blok masing-masing kondisi.
*/
package main

import "fmt"

func main()  {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n",percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
