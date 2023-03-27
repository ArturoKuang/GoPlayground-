package main

import (
	"fmt"
	outside "video"
	bar "video/test"
	"video/testPackage"
)

func main() {
	fmt.Println(testPackage.Great())
	fmt.Println("dsadsa")
	fmt.Print(bar.F())
	outside.A()
}
