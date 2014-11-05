package main

import "github.com/andyleap/modular/module"

type Test int

func (t *Test) Test(args *string, reply *string) error {
	*reply = *args + " World!"
	return nil
}

func main() {
	module := module.New()
	test := new(Test)
	module.Register(test)
	module.Serve()
}
