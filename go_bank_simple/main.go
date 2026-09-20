package main

import (
	"fmt"
)

func main() {

	choice := 0
	choiceUser := 0
	var p person
	var logado bool
	var login, password string

	fmt.Println("=======================================")
	for choice != 1 {

		fmt.Println("Digite 1 para sair do programa")
		fmt.Println("Digite 2 para criar sua conta")
		fmt.Println("Digite 3 para logar no banco")

		fmt.Scanln(&choice)
		if choice == 1 {
			return
		}
		if choice == 2 {
			fmt.Println("Digite seu nome")
			fmt.Scanln(&p.name)

			fmt.Println("Digite seu sobrenome")
			fmt.Scanln(&p.lastName)

			fmt.Println("Digite sua senha")
			fmt.Scanln(&p.password)

			dbPerson = append(dbPerson, p)
		}
		if choice == 3 {
			logado = false
			fmt.Println("Digite seu login")
			fmt.Scanln(&login)
			fmt.Println("Digite sua senha")
			fmt.Scanln(&password)
			for _, i := range dbPerson {
				if login == i.name { //verifi the login
					if password == i.password { //verific the passord

						fmt.Println("Logado com sucesso")
						fmt.Println("Digite 1 para sair do programa")
						fmt.Println("Digite 2 para ver seu extrato bancário")
						fmt.Println("Digite 3 para sair adicionar valor")
						fmt.Println("Digite 4 para fazer uma transferencia")
						fmt.Println("Digite 5 para ver os usuairos")
						logado = true
						choiceUser = 0
						for choiceUser != 1 {
							fmt.Scanln(&choiceUser)
							if choiceUser == 2 {
								result, err := showExtract(login)
								if err != nil {
									fmt.Println(err)
									return
								}
								fmt.Printf("seu extrato: %+v \n", result)
							}
							if choiceUser == 3 {
								fmt.Println("Digite o valor que você quer depositar")
								var value float32
								fmt.Scanln(&value)
								addValue(value, login)
							}
							if choiceUser == 4 {
								fmt.Println("Digite o nome da pessoa que você quer depositar o valor")
								var name string
								fmt.Scanln(&name)

								fmt.Println("Digite o valor que você quer depositar")
								var value float32
								fmt.Scanln(&value)

								transfValue(value, name, login)
							}
							if choiceUser == 5 {
								showUsers()
							}

							fmt.Println("Digite 1 para sair do programa")
							fmt.Println("Digite 2 para ver seu extrato bancário")
							fmt.Println("Digite 3 para sair adicionar valor")
							fmt.Println("Digite 4 para fazer uma transferencia")
							fmt.Println("Digite 5 para ver os usuairos")
						}

					}
				}
			}
			if logado == false {
				fmt.Println("Usuario não encontrado")
			}

		}
		fmt.Println("=======================================")

	}
}
