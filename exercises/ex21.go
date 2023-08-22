package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

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
	Execute() (string, error)
}

// A printing plugin: will print a message

type PrintPlugin struct {
	message string
}

func (p *PrintPlugin) Execute() (string, error) {
	return p.message, nil
}

// A math plugin: will perform a math operation (it needs 2 operands and 1 operator)

type MathPlugin struct {
	op1      decimal.Decimal // you can use float64/32
	op2      decimal.Decimal
	operator string
}

func (m *MathPlugin) Execute() (string, error) {
	var result decimal.Decimal
	switch m.operator {
	case "ADD":
		result = m.op1.Add(m.op2)
	case "DIV":
		result = m.op1.Div(m.op2)
	case "MUL":
		result = m.op1.Mul(m.op2)
	case "SUB":
		result = m.op1.Sub(m.op2)
	default:
		return "", fmt.Errorf("Operator not recognized: %s", m.operator)
	}

	return fmt.Sprintf(
		"%s %s %s = %s",
		m.op1.String(),
		m.operator,
		m.op2.String(),
		result.String(),
	), nil
}

// Commander pattern
type PluginCommander struct {
	plugins []Plugin
}

func (p *PluginCommander) Execute() {
	for _, plugin := range p.plugins {
		response, err := plugin.Execute()
		if err != nil {
			fmt.Println("Error when executing plugin", err)
			continue
		}

		fmt.Println(response)
	}
}

func main() {
	start := &PrintPlugin{message: "Starting plugins execution."}
	mathPlugin := &MathPlugin{
		op1:      decimal.NewFromFloat(3.5),
		op2:      decimal.NewFromFloat(2.0),
		operator: "MUL",
	}
	end := &PrintPlugin{message: "Plugins finished executing."}

	commander := &PluginCommander{
		plugins: []Plugin{start, mathPlugin, end},
	}

	commander.Execute()
}
