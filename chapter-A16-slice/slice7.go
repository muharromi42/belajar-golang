// A.16.7. Fungsi copy()
package main

import "fmt"

func main()  {
	dst := make([]string,3)
	src := []string{"nanas","pisang","jeyuk","apel"}
	n := copy(dst,src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}
