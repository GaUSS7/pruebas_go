package main

import "fmt"

func main() {
	/*  conversor de minutos a hora*/
	var minuto int
	fmt.Println("Ingrese los minutos que desee convertir: ")
	fmt.Scanln(&minuto)
	fmt.Println("Ingresaste: ", minuto, " minuto, la conversin a hora es de: ", minuto/60)
	var hora int
	fmt.Println("Ingrese las horas  que desee convertir: ")
	fmt.Scanln(&hora)
	fmt.Println("Ingresaste: ", hora, " horas, la conversin a minutos es de: ", hora*60)

}
