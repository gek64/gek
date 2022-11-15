package gError

import (
	"fmt"
)

func ExampleErr_Error() {
	var err Err
	err.New(404, "can't found")
	fmt.Print(err.Error())
	// Output:
	// code: 404
	// msg: can't found
}

func ExampleErr_New() {
	var err Err
	err.New(404, "can't found")
	fmt.Print(err)
	// Output:
	// {404 can't found}
}

func ExampleErr_Json() {
	var err Err
	err.New(404, "can't found")
	fmt.Print(err.Json())
	// Output:
	// {"Code":404,"Msg":"can't found"}
}
