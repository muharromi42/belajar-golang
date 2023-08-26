// A.16.2. Hubungan Slice Dengan Array & Operasi Slice
package main

import "fmt"

func main()  {
	var buah = []string {"apel","mangga","pisang"}
	var newBuah = buah[0:2]

	fmt.Println(newBuah)
}
