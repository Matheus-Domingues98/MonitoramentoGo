package main

import (
	c "OrientacaoObjeto/banco/contas"
	//cl "OrientacaoObjeto/banco/clientes"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDenis := c.ContaPoupanca{}
	contaDenis.Depositar(100)
	PagarBoleto(&contaDenis, 60)

	fmt.Println(contaDenis.ObterSaldo())

	contaLuisa := c.ContaCorrente{}
	contaLuisa.Depositar(500)
	PagarBoleto(&contaLuisa, 200)
	fmt.Println(contaLuisa.ObterSaldo())
}