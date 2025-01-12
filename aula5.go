package main

import "fmt"

func main() {
	var nome string = "André"
	var idade int = 32
	var altura float32 = 1.61
	var peso float32 = 84.5

	var versaoPrograma float32 = 2.2

	fmt.Println("Olá Sr.", nome, ", em nossos registros consta que você tem", idade, "anos, mede", altura, "de altura e pesa", peso, "kg.")
	fmt.Println("Este programa está na versão ", versaoPrograma)
}
