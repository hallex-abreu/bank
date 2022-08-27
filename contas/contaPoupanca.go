package contas

import (
	"fmt"
	"os"

	"github.com/hallex-abreu/bank/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Saque(valor float64) float64 {
	if valor > c.saldo {
		fmt.Println("saldo insuficiente")
		os.Exit(0)
	}
	return c.saldo - valor
}

func (c *ContaPoupanca) Deposito(valor float64) float64 {
	if valor < 0 {
		fmt.Println("Valor não pode ser negativo")
		return c.saldo
	} else {
		c.saldo += valor
		return c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
