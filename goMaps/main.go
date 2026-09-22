package main

import (
	"fmt"
)

func main() {
	//var colors map[int]string

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"green": "#4bf746",
	}

	//colors := make(map[int]string)

	//colors[10] = "#ffffff"

	//delete(colors, 10)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("key: %+v value:: %+v \n", color, hex)
	}
}
