package main

import (
	"fmt"
)

func main() { // HLmain
	a := &audience{"BigSky DevCon"} // HLmain
	fmt.Printf("Hello %s!\n", a)    // HLmain
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
