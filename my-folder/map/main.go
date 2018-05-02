package main

import "fmt"

func main() {

	// 1st way to declare a map
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "FFFFFF",
	}

	// empty map declaration #1
	// var colors map[string]string

	// empty map declaration #2
	// colors := make(map[string]string)

	// add value to map
	// colors["white"] = "#FFFFFF"

	//delete value from map
	// delete(colors, "white")

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
