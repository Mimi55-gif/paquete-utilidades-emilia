package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	opc := 0
	var dol float64 = 0.0
	//Bucle
	for {
		fmt.Println("Selecciona la moneda a tranformar")
		fmt.Println("1.Euro\n2.Libra\n3.Won\n4.BTC")
		fmt.Scan(&opc)
		//Condición para poder salir del programa
		if opc != 0 {
			switch opc {

			case 1:
				fmt.Println("Por favor ingrese la cantidad de dolares que desea convertir")
				fmt.Scan(&dol)
				//Operación por moneda
				euro := dol * 0.86
				//Impresión con la variable de asignada a cada moneda
				fmt.Println("dolares = $", dol, "\nEuros = €", euro)
			case 2:
				fmt.Println("Por favor ingrese la cantidad de dolares que desea convertir")
				fmt.Scan(&dol)
				libras := dol * 0.75
				fmt.Println("dolares = $", dol, "\nLibras = £", libras)
			case 3:
				fmt.Println("Por favor ingrese la cantidad de dolares que desea convertir")
				fmt.Scan(&dol)
				won := dol * 1411.17
				fmt.Println("dolares = $", dol, "\nWones = ₩", won)
			case 4:
				fmt.Println("Por favor ingrese la cantidad de dolares que desea convertir")
				fmt.Scan(&dol)
				BTC := dol * 0.0000091
				fmt.Printf("dolares = $ %.2f\nBTC = ₿ %f \n", dol, BTC)
			}

		} else {
			//En caso de aplastar 0 sale del programa
			fmt.Println("Gracias por user el convertidor\n Gracias por todo profe Ronie")
			break
		}

	}

	lect := bufio.NewReader(os.Stdin)
	a, e, i, o, u := 0, 0, 0, 0, 0

	fmt.Println("Contador de vocales.")
	println("Ingresa una frase")
	fra, er := lect.ReadString('\n')
	fra = strings.ToLower(fra)

	if er != nil {
		fmt.Print("Ha ocurrido un error al leer la frase, adios.")
		return
	} else {
		for _, carac := range fra {
			switch carac {
			case 'a':
				a++

			case 'e':
				e++
			case 'i':
				i++
			case 'o':
				o++
			case 'u':
				u++
			}

		}
		fmt.Print("\nCantidad de Vocales en la frase: ", fra)
		fmt.Printf("a: %d\n", a)
		fmt.Printf("e: %d\n", e)
		fmt.Printf("i: %d\n", i)
		fmt.Printf("o: %d\n", o)
		fmt.Printf("u: %d\n", u)

	}

}
