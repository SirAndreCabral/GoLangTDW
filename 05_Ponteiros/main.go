package main

import "fmt"

func main() {

	valor := 5
	ponteiroDoValor := &valor

	fmt.Println("Valor", valor)
	fmt.Println("Endereco de memoria Valor", &valor)
	fmt.Println("Endereco de memoria Valor", *ponteiroDoValor)
}
