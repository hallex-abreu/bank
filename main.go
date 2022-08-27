package main

import (
	"fmt"

	"github.com/hallex-abreu/bank/clientes"
	"github.com/hallex-abreu/bank/contas"
)

func PagarBoleto(conta verificaConta, valorBoleto float64) {
	conta.Saque(valorBoleto)
}

type verificaConta interface {
	Saque(valor float64) float64
}

func main() {
	contaHallex := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Hallex",
			Cpf:       "06391271348",
			Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456}

	contaHallex.Deposito(200)

	contaLaiza := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Laiza",
			Cpf:       "07201316346",
			Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456}

	contaLaiza.Deposito(100)

	status := contaHallex.Trasnferir(100, &contaLaiza)
	if status {
		fmt.Println("Transferencia executada com sucesso")
	} else {
		fmt.Println("Transferencia não foi executada")
	}

	PagarBoleto(&contaHallex, 100)

	fmt.Println("Hallex Corrente:", contaHallex.ObterSaldo())
	fmt.Println("Laiza Poupança", contaLaiza.ObterSaldo())
}
