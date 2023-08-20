// A.15.1. Pengisian Elemen Array yang Melebihi Alokasi Awal
package main

import "fmt"

func main()  {
	var names [4] string
	names[0] = "trafalgar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"
	names[4] = "ez" //kode ini akan menghasilkan error

}
