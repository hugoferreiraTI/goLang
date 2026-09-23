package main

import (
	"errors"
	"fmt"
)

type metodoDePagamento interface {
	pagar(valor float64) error
}

type pix struct {
	valor float64
}

type cartaoDeCredito struct {
	limite float64
	uso    float64
}

type carrinho struct {
	valor       float64
	nameProduto string
}

func main() {
	carrinho := []carrinho{{valor: 200, nameProduto: "Sapat"}, {valor: 500, nameProduto: "controle"}}

	pixValor := pix{
		valor: 2000,
	}

	cartao := cartaoDeCredito{
		limite: 1000,
	}

	finalizarCompra(carrinho, pixValor)
	finalizarCompra(carrinho, cartao)

}

func finalizarCompra(items []carrinho, metodo metodoDePagamento) error {
	var valorFinal float64
	for _, valorProduto := range items {
		valorFinal += valorProduto.valor
	}

	err := metodo.pagar(valorFinal)
	if err != nil {
		return err
	}

	return nil
}

func (p pix) pagar(valor float64) error {
	if p.valor <= valor {
		return errors.New("Valor insuficiente")
	}
	p.valor = p.valor - valor

	fmt.Println("Foi seu gay ai o valor novo: ", p.valor)

	return nil
}

func (c cartaoDeCredito) pagar(valor float64) error {
	if c.limite < valor {
		return errors.New("Limite insuficiente")
	}
	c.limite -= valor

	fmt.Println("Foi seu gay ai o valor novo: ", c.limite)
	return nil
}
