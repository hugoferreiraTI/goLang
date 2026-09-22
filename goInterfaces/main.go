package main

import "fmt"

type bot interface {
	getGreeting() string
}

//type bot interface {
// getGreeting(string, int)(string, error)
//} This a exempl for others interfaces, this case use two arguments and return two values - The others function contract is true

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generatin an english gretting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//VERY custom logic for generatin an english gretting//
	return "Hi teste!"
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sp spanishBot) {
//	fmt.Println(sp.getGrreting())
//}
