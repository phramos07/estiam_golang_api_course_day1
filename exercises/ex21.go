package main

/*
Exercise: Implementing a Plugin System

Create a simple plugin system that allows you to dynamically load and use different plugins. Plugins can provide various functionalities, and you'll use interfaces to define a common contract for them.

Define a plugin interface Plugin:
	Create an interface named Plugin that declares a method for executing the plugin's functionality.

Implement plugin structures:
	Create multiple structs that represent different plugins.
	Each struct should implement the Plugin interface with its own Execute method.

Create a function to use plugins:
	Create a function that takes a Plugin interface as a parameter and executes its functionality.

Implement a simple plugin loader:
	Create a map where you can store plugin names and corresponding plugin instances.
	Load a few instances of your plugin structs into the map.

Test on main.
*/

// Plugin interface: defines the contract for every plugin.
type Plugin interface {
	Execute() string
}

// A printing plugin: will print a message

type PrintPlugin struct{}

// A math plugin: will perform a math operation (it needs 2 operands and 1 operator)

type MathPlugin struct{}

func main() {

}
