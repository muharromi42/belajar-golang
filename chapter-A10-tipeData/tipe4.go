// A.10.4. Tipe Data string
package main

import "fmt"

func main()  {
	var message string = "halo"
	var message2 string = `Nama saya "romi"
	belajar golang sangat mudah
	salam kenal.`

	fmt.Printf("message : %s \n", message)
	fmt.Printf("message : %s \n", message2)
}
