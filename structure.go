package main

import "fmt"

type Student struct{
	roll int
	name string
}

func main() {
	s1 := Student{12, "Akhil"}
	fmt.Println(s1)
}
