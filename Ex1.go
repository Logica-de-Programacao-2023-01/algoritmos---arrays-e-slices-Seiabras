// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
package main

import "fmt"

func main() {
	//criar array
	array := [3]int{1, 2, 3}
	//calculo soma
	soma := 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}
	//calcular valor soma
	fmt.Print("Esse é o valor da soma: ", soma)
}
