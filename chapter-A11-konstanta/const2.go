// A.11.2. Deklarasi Multi Konstanta
package main

import "fmt"

func main()  {
	const(
		square = "kotak"
		isToday bool = true
		numeric uint8 = 1
		angka
		floatNum = 2.2
	)
	/*
    Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai
    yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.
*/

	const satu,dua = 1,2
	const tiga,empat string = "tiga", "empat"

	fmt.Println("hai",square,isToday,numeric,angka,floatNum,satu,dua,tiga,empat)
}
