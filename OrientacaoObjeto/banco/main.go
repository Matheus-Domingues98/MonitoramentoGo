package main

import (
	c "OrientacaoObjeto/banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := c.ContaCorrente{
		Titular: "Silvia",
		Saldo:   1000.0,
	}

	contaDoGustavo := c.ContaCorrente{
		Titular: "Gustavo",
		Saldo:   500.0,
	}

	status := contaDaSilvia.Transferir(900.0, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
