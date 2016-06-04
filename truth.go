package main

import (
	"fmt"
)

func main() {
	fun := true
	youAreHaving := 0
	// 0, "", nil, false, ...
	if youAreHaving == fun {
		fmt.Println("Yaw!")
	} else {
		fmt.Println("Damn")
	}
}
