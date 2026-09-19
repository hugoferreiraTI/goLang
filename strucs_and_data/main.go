package main

import "fmt"

type contactInf struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInf
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//var alex person
	//alex.firstName = "alex"
	//alex.lastName =  "anderson"
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName:  "Hugo",
		lastName:   "Ferreira",
		contactInf: contactInf{email: "TESTE@", zipCode: 123123},
	}

	//jimPointer := &jim

	jim.updateName("jimmy2ww2")

	jim.printFunc()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
} 

func (p person) printFunc() {
	fmt.Printf("%+v, %+v", p.firstName, p.contactInf.zipCode)
}
