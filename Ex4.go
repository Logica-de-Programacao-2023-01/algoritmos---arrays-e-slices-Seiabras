// Crie um Slice de inteiros e solicite ao usuário que informe
// o tamanho do Slice e os valores dos elementos. Em seguida,
// imprima o Slice e a soma dos valores armazenados.
package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scan(&tamanho)

	lista := make([]int, tamanho)

	//for i := range slice{
	//fmt.Printf("Digite o valor para o elemento %d: ", i)
	//fmt.Scanln(&slice[i])
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i)
		fmt.Scan(&lista[i])
	}

	fmt.Println("O slice informado é:", lista)

	soma := 0
	for _, valor := range lista {
		soma += valor
	}
	fmt.Println("A soma dos valores armazendos é", soma)

}
