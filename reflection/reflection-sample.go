package main

import (
	"fmt"
	"reflect"
)

type MyClient struct {
	ClientArg string
}

func (client *MyClient) CreateMethod1(argument string, arg2 int) string {
	fmt.Printf("Running create method 1 %s and %d\n", argument, arg2)
	return "value returned here"
}

func (client *MyClient) CreateMethod2() {
	fmt.Println("running create Method 2 with arg " + client.ClientArg)
}

func main() {
	var client MyClient

	// Equivalent of calling : client.CreateMethod1("hitesh")
	args := make([]reflect.Value, 2)
	args[0] = reflect.ValueOf("Hitesh")
	args[1] = reflect.ValueOf(1)

	value := reflect.ValueOf(&client).MethodByName("CreateMethod1").Call(args)
	fmt.Printf("Value : %s\n", value)
	// ----

	// Equivalient of calling client.CreateMethod2()
	value = reflect.ValueOf(&client).MethodByName("CreateMethod2").Call([]reflect.Value{})
	// ---
}
