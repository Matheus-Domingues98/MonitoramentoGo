package contas

import (
	cl "OrientacaoObjeto/banco/clientes"
)

type ContaPoupanca struct {
	Titular cl.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}
	return "saldo insuficiente"
}

//Função com 2 retornos, uma string e um float64
func (c *ContaPoupanca) Depositar (valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso", c.saldo
	}
	return "Valor inválido para depósito", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}