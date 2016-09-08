package main

import "fmt"

import (
	calc_lib "calculater/calc_lib/calc_lib_01"
)

func main() {
	//var b bool
	fmt.Printf("Hello, world; or lljflwei or 你好世界\n")
	b := calc_lib.I_divide_2(3)
	fmt.Printf("calc_lib.I_divide_2(3): %t\n", b)
	b = calc_lib.Get_temprature_of_fire(9)
	fmt.Printf("calc_lib.Get_temprature_of_fire(9): %t\n", b)
}
