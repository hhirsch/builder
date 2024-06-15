package main

func main() {
	var interpreter Interpreter = *NewInterpreter()
	interpreter.load("setup.bld")
}
