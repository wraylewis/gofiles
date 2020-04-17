package main

import "fmt"

func main() {
	s := "Hello playground"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

}
