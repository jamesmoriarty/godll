package main

import (
	"C" // https://github.com/golang/go/issues/22192
	"fmt"
	"syscall"
)

var (
	dll  = syscall.MustLoadDLL("example.dll")
	proc = dll.MustFindProc("Sum")
)

func main() {
	fmt.Printf("dll=%v proc=%v\n", dll, proc)

	r1, _, err := proc.Call(1, 2)

	if err.Error() != "The operation completed successfully." {
		fmt.Printf("%v", err)
	}

	fmt.Printf("1 + 2 = %v\n", r1)
}
