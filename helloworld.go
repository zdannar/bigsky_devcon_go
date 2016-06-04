package main

import (
	"fmt"
)

// STRUCT OMIT
type audience struct {
	statement string
}

func (a *audience) String() string {
	return a.statement
}

// STRUCTE OMIT

func main() { // HLmain
	a := &audience{"BigSky DevCon"} // HLmain
	fmt.Printf("Hello %s!\n", a)    // HLmain
} // HLmain
