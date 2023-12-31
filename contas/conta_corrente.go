package contas

import "github.com/leandroyyy/go-poo/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo = c.saldo - valorSaque
		return "saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "deposito realizado com sucesso", c.saldo
	}

	return "o valor do deposito menor que zero", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
