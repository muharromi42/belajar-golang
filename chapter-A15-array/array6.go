// A.15.5. Array Multidimensi
package main

import "fmt"

func main()  {
	var angka1 = [2][3] int {[3]int {10,11,12}, [3]int{13,14,15}}
	var angka2 = [2][3] int {{16,17,18},{19,20,21}}

	fmt.Println("angka1 :", angka1)
	fmt.Println("angka2 :", angka2)
}
