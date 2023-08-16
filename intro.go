package main

import "fmt"

func main() {
	fmt.Print("Hola Mundo Golang ! ")

	var edad = 32
	var nombre = "Javier"

	fmt.Print("mi nombres es ", nombre, " y tengo ", edad, " años")

	//Entradas
	var web string
	fmt.Println("Introduce una entrada")
	fmt.Scan(&web)
	fmt.Println("Tu entrada fue: ", web)

	//Condicionales

	var altura int

	fmt.Print("¿Cual es tu altura en cms?")
	fmt.Scan(&altura)

	if altura >= 180 {
		fmt.Println("Eres Alto ! felicidades")
	} else {
		fmt.Println("Eres bajo que pena")
	}

	//Funciones

	fmt.Println("Calculando IMC con peso: ", calcularIMC(2))
	fmt.Println("------------------------------------------")
	fmt.Println("Calculando IMC con Altura: ", calcularIMCAltura(altura))

}

func calcularIMC(peso int) int {

	var resultado int
	resultado = peso * 5
	return resultado
}

func calcularIMCAltura(altura int) int {
	var resultado int
	resultado = altura * 3
	return resultado
}
