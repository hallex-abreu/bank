package contas

import (
	"fmt"
	"os"

	"github.com/hallex-abreu/bank/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Saque(valor float64) float64 {
	if valor > c.saldo {
		fmt.Println("saldo insuficiente")
		os.Exit(0)
	}
	return c.saldo - valor
}

func (c *ContaCorrente) Deposito(valor float64) float64 {
	if valor < 0 {
		fmt.Println("Valor não pode ser negativo")
		return c.saldo
	} else {
		c.saldo += valor
		return c.saldo
	}
}

func (c *ContaCorrente) Trasnferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor <= c.saldo {
		c.saldo -= valor
		contaDestino.saldo = contaDestino.Deposito(valor)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
