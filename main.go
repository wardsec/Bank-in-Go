package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(verificarConta, valorDoBoleto float64) {
	conta.Sacar
}
func main() {
	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())
}
