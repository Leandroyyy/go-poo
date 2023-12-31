package main

import (
	"fmt"

	"github.com/leandroyyy/go-poo/clientes"
	"github.com/leandroyyy/go-poo/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaBruno := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "bruno",
			Cpf:       "12312321",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 50,
		NumeroConta:   1234,
	}

	contaBruno.Depositar(200)

	PagarBoleto(&contaBruno, 40)

	fmt.Println(contaBruno.ObterSaldo())
}
