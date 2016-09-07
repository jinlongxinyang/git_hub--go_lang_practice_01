package main

import "fmt"
import "even"

func main() {
	//var b bool
	fmt.Printf("Hello, world; or lljflwei or 你好世界\n")
	b := even.Even(3)
	fmt.Printf("b: %t\n", b)
}
