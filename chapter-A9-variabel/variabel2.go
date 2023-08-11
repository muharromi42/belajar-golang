//A.9.2. Deklarasi Variabel Menggunakan Keyword var
package main

import "fmt"

func main () {
	var firstname string = "Romay"

	var lastname string
	lastname = "Ardhana"

	fmt.Printf("halo Romay Ardhana \n")
	fmt.Printf("halo %s %s \n",firstname,lastname)
	fmt.Println("halo", firstname,lastname + "!")
}
