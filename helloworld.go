package main

import (
	"fmt"
)

// Hey look, there is a main function and I can put it wherever // HLmain
// I want for easy reading! // HLmain
func main() { // HLmain
	a := &audience{"BigSky DevCon"}
	fmt.Printf("Hello %s!\n", a)
} // HLmain

// STRUCT OMIT
type audience struct {
	statement string
}

// Method of type audience
func (a *audience) String() string {
	return a.statement
}

// STRUCTE OMIT
