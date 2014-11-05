package main

import (
	"github.com/andyleap/modular/loader"
	"fmt"
)

func main() {
	module, _ := loader.Load("../module/module")
	arg := "test"
	var reply string
	module.Call("Test.Test", arg, &reply)
	fmt.Println(reply)
	module.Close()
}
