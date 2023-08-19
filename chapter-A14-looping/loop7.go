// A.14.7. Pemanfaatan Label Dalam Perulangan
package main

import "fmt"

func main()  {
	outerloop:
	for i := 0; i<5; i++ {
		for j :=0; j<5; j++ {
			if i == 3 {
				break outerloop
			}

			fmt.Println("matriks [",i,"][",j,"]", "\n")
		}
	}
}
