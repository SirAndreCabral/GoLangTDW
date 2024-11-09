package main

import "fmt"

type PessoaID int

type Pessoa struct {
	PessoaID  PessoaID
	Nome      string
	Sobrenome string
	Idade     int
}

type Produto struct {
	Nome   string
	Status string
}

//Letra minuscula = privada
func (p Pessoa) saudar() {
	fmt.Println("Olá, ", p.Nome)
}

func (p Pessoa) SaudarMetodoPublico() {
	fmt.Println("Olá, ", p.Nome)
}

func (p Produto) IsActive() bool {
	return p.Status == "Active"
}

func main() {
	pessoa := Pessoa{
		PessoaID:  1,
		Nome:      "Maria",
		Sobrenome: "Silva",
		Idade:     20,
	}

	pessoa.saudar()

	fmt.Println(pessoa)

	fmt.Println("Nome da pessoa", pessoa.Nome)
}
