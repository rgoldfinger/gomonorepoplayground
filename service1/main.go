package main

import (
	"fmt"

	"example.com/monorepo/package1"
)


func main() {
	fmt.Println("service 1")
	fmt.Println(package1.RevString("fooooo"))
}
