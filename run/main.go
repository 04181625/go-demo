package main

func foo(arg_val int) *int {

	var foo_val int = 11
	return &foo_val
}

func main() {

	main_val := foo(666)

	println(*main_val)
}
