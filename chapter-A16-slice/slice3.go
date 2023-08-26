// A.16.3. Slice Merupakan Tipe Data Reference
/*rtinya jika ada slice baru yang terbentuk dari slice lama, maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
*/
package main

import "fmt"

func main()  {
	var buah = []string{"apel","anggur","pisang","melon"}

	var aBuah = buah[0:3]
	var bBuah = buah[1:4]

	var aaBuah = aBuah[1:2]
	var baBuah = bBuah[0:1]

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(baBuah)

	// buah anggur di ubah menjadi nanas
	baBuah[0] = "nanas"

	fmt.Println(buah)
	fmt.Println(aBuah)
	fmt.Println(bBuah)
	fmt.Println(aaBuah)
	fmt.Println(baBuah)
}
