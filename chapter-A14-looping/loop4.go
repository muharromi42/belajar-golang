// A.14.4. Penggunaan Keyword for - range
package main

import "fmt"

func main()  {
	var xs = "123" //string
	for i,v := range xs {
		fmt.Println("index=",i, "Value=",v)
	}

	var ys = [5]int{10,20,30,40,50} //array
	for _, v := range ys {
		fmt.Println("Values=",v)
	}

	var zs = ys[0:2] //slice
	for _, v := range zs {
		fmt.Println("value=",v)
	}

	var kvs = map[byte]int{'a':0, 'b':1, 'c':2} //map
	for k,v := range kvs {
		fmt.Println("Keys=", k, "values=",v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
		fmt.Println("done")
	}
}
