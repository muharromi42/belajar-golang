// A.9.3. Deklarasi Variabel Tanpa Tipe Data
package main

import "fmt"


func main () {
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstname string = "Romay"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastname := "Ardhana"

	fmt.Printf("Halo %s %s \n", firstname, lastname)
}
