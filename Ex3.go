// Crie um Array de floats com 4 elementos e calcule o produto
// dos valores armazenados no Array.
package main

import "fmt"

func main() {
	lista := [4]float64{1.1, 1.3, 1.4, 1.5}
	produto := 1.0

	for i := 0; i < len(lista); i++ {
		produto *= lista[i]
	}
	fmt.Println("O valor do produto é:", produto)

}
