package contas

import "github.com/leandroyyy/go-poo/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo = c.saldo - valorSaque
		return "saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "deposito realizado com sucesso", c.saldo
	}

	return "o valor do deposito menor que zero", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
