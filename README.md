** Bank in Golang **
* este script tem como funções basicas: 
    - Adicionar saldo em contas.
    - Transferir Saldos em contas.
    - Remover saldo em contas.

* Como rodar este projeto ?
Apenas utilize o comando: go run main.go

Desta forma conseguimos depositar a visualizar o saldo:
``````` 
func main() {
    contaExemplo := contas.ContaCorrente{}
    contaExemplo.Depositar(100)

    fmt.Println(contaExemplo.ObterSaldo())
}
````````
