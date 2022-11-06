//criar um programa em GO que trabalhe com o operdaor % e for.
//Desafio:
//Vocês e seus colegas querem criar um código que exiba os números
//entre 1 e 100 que são divisíveis por 3.

package main

import "fmt"

func main(){
	for i:=1; i<=100; i++{
		if i%3 == 0 {
			fmt.Println("é divisível por 3:",i)
		} else {
			fmt.Println("não é divisível por 3:",i)
		}
	}
}