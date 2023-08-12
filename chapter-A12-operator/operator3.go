// A.12.3. Operator Logika
package main

import "fmt"

func main() {
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left and right \t(%t) \n",leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left or right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("left \t\t(%t) \n", leftReverse)
}
