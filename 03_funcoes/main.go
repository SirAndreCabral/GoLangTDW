package main

import "fmt"

func main() {

	// fmt.Println(somaDoisNumeros(2, 2))
	result, resultString := somaDoisNumeros(2, 2)
	fmt.Println(fmt.Sprintf("%s - %d", resultString, result))
}

func somaDoisNumeros(numero1 int, numero2 int) (int, string) {
	return (numero1 + numero2), "A soma dos dois numeros é"
}
