package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	fmt.Println("the first init method")
}

func init() {
	fmt.Println("the second init method")
}

func init() {
	fmt.Println("the init method for reading flag")
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

/*
	init函数先于main函数执行。每个包可以有多个init函数，每个init函数执行顺序没确定。
*/
func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
