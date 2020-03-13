/*
Package pathfinder efectua el routing en una topología de Hipercubo.
*/
package pathfinder

import "strconv"

/*PathFinder searchs for the path from a node origen to destino in a dimension hipercube*/
func PathFinder(origen, destino int64, dimension int) string {
	var mask int64
	camino := make([]int64, 0)
	ret := ""

	// Agrego el primer nodo que es el origen
	camino = append(camino, origen)
	//Calculo la distancia
	dist := origen ^ destino

	for i := 0; i < dimension; i++ {
		// Cargo el valor de la mascara en su slice
		mask = 1 << uint8(i)
		// Si tengo que cambiar el bit
		if step := dist & mask; step != 0 {
			// Variable a devolver
			var node int64
			//Lo cambio
			if j := len(camino); j == 1 {
				node = origen ^ step
			} else {
				node = camino[j-1] ^ step
			}
			camino = append(camino, node)
		}
	}

	for i := range camino {
		nodeS := strconv.FormatInt(camino[i], 2)
		for len(nodeS) < dimension {
			nodeS = "0" + nodeS
		}
		ret = ret + nodeS
		if i < len(camino)-1 {
			ret = ret + " -> "
		}
	}

	return ret
}