package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			defer fmt.Println("C.defer")
		}()
	}()
}
