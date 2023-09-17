package main

import "plugin"

func main() {
	plugin, _ := plugin.Open("plugin.so")

	helloWorldSymbol, _ := plugin.Lookup("HelloWorld")

	helloWorld := helloWorldSymbol.(func())

	helloWorld()
}
