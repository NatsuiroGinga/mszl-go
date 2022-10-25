package main

import "fmt"

func main() {

	num := 10
	fmt.Println("==first==")
	modifyNum1(num)
	fmt.Println("old num = ", num)

	fmt.Println("==second==")
	modifyNum2(&num)
	fmt.Println("old num = ", num)
}

func modifyNum1(num int) {
	num = 20
	fmt.Println("new num = ", num)
}

func modifyNum2(num *int) {
	*num = 30
	fmt.Println("new num = ", *num)
}
