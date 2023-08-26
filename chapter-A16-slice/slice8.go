// A.16.8. Pengaksesan Elemen Slice Dengan 3 Indeks
// 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya.
package main

import "fmt"

func main()  {
	var buah = []string{"apel","nanas","pisang"}
	var aBuah = buah[0:2]
	var bBuah = buah[0:2:2]

	fmt.Println(buah)      // ["apple", "grape", "banana"]
	fmt.Println(len(buah)) // len: 3
	fmt.Println(cap(buah)) // cap: 3

	fmt.Println(aBuah)      // ["apple", "grape"]
	fmt.Println(len(aBuah)) // len: 2
	fmt.Println(cap(aBuah)) // cap: 3

	fmt.Println(bBuah)      // ["apple", "grape"]
	fmt.Println(len(bBuah)) // len: 2
	fmt.Println(cap(bBuah)) // cap: 2
}
