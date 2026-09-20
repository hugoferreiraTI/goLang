package main

type person struct {
	name     string
	lastName string
	password string
	extract  float32
}

var dbPerson = []person{
	{name: "Ana", lastName: "Silva", password: "senhaAna123", extract: 1540.50},
	{name: "Carlos", lastName: "Oliveira", password: "carlos@pass", extract: 250.00},
	{name: "Beatriz", lastName: "Souza", password: "bia#2024", extract: 4890.75},
	{name: "Lucas", lastName: "Pereira", password: "lucas_pwd", extract: 15.20},
	{name: "Mariana", lastName: "Costa", password: "mariCost@!", extract: 950.10},
	{name: "Gabriel", lastName: "Santos", password: "gabriel_pass", extract: 12500.00},
	{name: "Juliana", lastName: "Almeida", password: "ju#senha789", extract: 340.55},
	{name: "Rafael", lastName: "Lima", password: "rafael2023", extract: 55.00},
	{name: "Fernanda", lastName: "Rocha", password: "fer_rocha_!", extract: 2100.80},
	{name: "Rodrigo", lastName: "Martins", password: "rodrigo!99", extract: 0.00},
}
