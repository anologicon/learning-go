package main

import "fmt"

/*
 * Instruções de retorno limpas devem ser usadas apenas em funções curtas, 
 * como no exemplo mostrado aqui. Elas podem prejudicar 
 * a legibilidade em funções mais longas.
 */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(25))
}
