package main

import "fmt"

func main() {
	// Valor do ponto de ebulição da água em Kelvin
	const kelvin float64 = 373.15

	// Conversão de Kelvin para Celsius
	celsius := kelvin - 273.15

	fmt.Printf("O ponto de ebulição da água em Celsius é: %.2f°C\n", celsius)
}
