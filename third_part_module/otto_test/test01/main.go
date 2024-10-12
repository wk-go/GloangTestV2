package main

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

func main() {
	vm := otto.New()

	// Evaluate a script
	vm.Run(`
    abc = 2 + 2;
    console.log("The value of abc is " + abc); // 4
`)
	// Get the value of abc
	if value, err := vm.Get("abc"); err == nil {
		if value_int, err := value.ToInteger(); err == nil {
			fmt.Printf("The value of abc in go is %#v, %#v\n", value_int, err)
		}
	}

	// Set a number
	vm.Set("def", 11)
	vm.Run(`
    console.log("The value of def is " + def);
    // The value of def is 11
`)
	// Set a string
	vm.Set("xyzzy", "Nothing happens.")
	vm.Run(`
	console.log("The value of xyzzy is ", "\""+xyzzy+"\"")
    console.log("The length of xyzzy is ", xyzzy.length); // 16
`)
	// Get the value of an expression
	value, _ := vm.Run("xyzzy.length")
	{
		// value is an int64 with a value of 16
		_value, _ := value.ToInteger()
		fmt.Printf("The value of xyzzy.length %#v\n", _value)
	}

	// An error happens
	v, err := vm.Run("abcdefghijlmnopqrstuvwxyz.length")
	if err != nil {
		// err = ReferenceError: abcdefghijlmnopqrstuvwxyz is not defined
		// If there is an error, then value.IsUndefined() is true
		fmt.Printf("An error happens: \"%s\", The value: %#v, value.IsUndefined(): %#v\n", err, v, v.IsUndefined())
	}

	// Set a Go function that returns something useful
	vm.Set("twoPlus", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToInteger()
		result, _ := vm.ToValue(2 + right)
		return result
	})
	vm.Set("sayHello", func(call otto.FunctionCall) otto.Value {
		right, _ := call.Argument(0).ToString()
		fmt.Printf("Hello, %s.\n", right)
		result, _ := vm.ToValue(nil)
		return result
	})

	// Use the functions in JavaScript
	result, _ := vm.Run(`
    sayHello("Xyzzy");      // Hello, Xyzzy.
    sayHello();             // Hello, undefined

    result = twoPlus(2.0); // 4
`)
	fmt.Printf("The result of twoPlus(2.0) is %#v\n", result)
}
