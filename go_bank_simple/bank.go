package main

import (
	"errors"
	"fmt"
)

func showExtract(name string) (float32, error) {
	for i := range dbPerson {
		if dbPerson[i].name == name {
			return dbPerson[i].extract, nil
		}
	}
	return 0, errors.New("Pessoa não encontrada")
}

func addValue(value float32, name string) error {
	if value <= 0 {
		return errors.New("Valor insuficiente")
	}
	for i := range dbPerson {
		if dbPerson[i].name == name {
			dbPerson[i].extract += value
			return nil
		}
	}
	return errors.New("Usuário não encontrado")
}

func transfValue(value float32, name string, nameUser string) error {
	if value <= 0 {
		return errors.New("Valor insuficiente")
	}
	for i := range dbPerson {
		if dbPerson[i].name == name {
			dbPerson[i].extract += value
			debitExtract(nameUser, value)
			return nil
		}
	}
	return errors.New("Pessoa não encontrada!")
}

func showUsers() error {
	if len(dbPerson) == 0 {
		return errors.New("Sem usuários no banco")
	}

	for i := range dbPerson {
		fmt.Printf("User: %+v, Extract: %+v \n", dbPerson[i].name, dbPerson[i].extract)
	}
	return nil
}

func debitExtract(nameUser string, value float32) error {
	if value <= 0 {
		return errors.New("Saldo insuficiente")
	}
	for i := range dbPerson {
		if dbPerson[i].name == nameUser {
			dbPerson[i].extract -= value
			return nil
		}
	}
	return errors.New("Usuario não encontrado")
}
