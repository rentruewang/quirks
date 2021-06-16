package main

import "fmt"

func deferFunc() {
	fmt.Println("defer function called")
}

func returnFunc() int {
	fmt.Println("return function called")
	return 0
}

func Function() int {
	defer deferFunc()
	return returnFunc()
}

// For the behavioral difference between named and unnamed return,
// see this link: https://stackoverflow.com/a/37249043

func unnamedFunc() string {
	// This is not javascript!
	str := "undefined"
	fmt.Println(str)
	defer func() {
		fmt.Println("defer", str)
		str = "defer"
		fmt.Println("defer", str)
	}()

	return func() string {
		fmt.Println("return", str)
		str = "return"
		return str
	}()
}

func namedFunc() (str string) {
	// This is not javascript!
	str = "undefined"
	fmt.Println(str)
	defer func() {
		fmt.Println("defer", str)
		str = "defer"
		fmt.Println("defer", str)
	}()

	return func() string {
		fmt.Println("return", str)
		str = "return"
		return str
	}()
}

func main() {
	Function()
	fmt.Println()
	fmt.Println(unnamedFunc())

	fmt.Println()
	fmt.Println(namedFunc())
}
