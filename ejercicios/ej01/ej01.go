package ej01

import "fmt"
// Camino de costo mínimo
// Dada una matriz con valores, cada celda contiene uno que representa el costo de pasar por esa celda.
// Se desea encontrar un camino que conecte la celda de la esquina superior izquierda con la esquina inferior derecha,
//cuyo coste total es mínimo. Devolver el costo del camino.

// **Ejemplo**
// ```
// matriz = 3,  2, 12, 15, 10
//          6, 19,  7, 11, 17
//          8,  5, 12, 32, 21
//          3, 20,  2,  9,  7

// output = 52
// ```

// **Resolver utilizando estrategia top-down.**

func CostoCaminoMinimo(matriz [][]int) int {
	return costoCaminoMinimo(matriz, 0, 0, make(map[string]int)) // el mapa va a guardar el valor minimo
}


func costoCaminoMinimo(matriz [][]int, i,j int, cache map[string]int) int {
	key := fmt.Sprintf("%d,%d",i,j)
	if result,ok := cache[key];ok{
		return result
	}

}