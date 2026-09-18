package main

import "fmt"

func EvenOrOdd(numbers []int) {
	for _, i := range numbers {
		if i%2 != 0 {
			fmt.Println("This number is a odd", i)
		}else {
		fmt.Println("This number is a even", i)
		}
	}
}
