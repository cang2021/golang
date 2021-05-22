package main

import "fmt"

// say Hi to someone
func SayHi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func main() {
	fmt.Println(SayHi("rob pike"))
}