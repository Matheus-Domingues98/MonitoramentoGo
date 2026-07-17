package main

import (
	c "OrientacaoObjeto/banco/contas"
	//cl "OrientacaoObjeto/banco/clientes"
	"fmt"
)

func main() {
	contaDenis := c.ContaPoupanca{}
	contaDenis.Depositar(100)
	contaDenis.Sacar(55)
	fmt.Println(contaDenis.ObterSaldo())
	contaPatricia := c.ContaCorrente{}
	fmt.Println(contaPatricia)

}