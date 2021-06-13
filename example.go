package main

import (
	"C" // https://github.com/golang/go/issues/22192
)

//export Sum
func Sum(arg1, arg2 int32) int32 {
	return arg1 + arg2
}

func main() {

}
