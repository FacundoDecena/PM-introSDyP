package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	continuar := true
	var origen, destino, dimension, maxDim int
	reader := bufio.NewReader(os.Stdin)
	for continuar {
		CallClear()
		fmt.Println("Ingrese una dimension: ")
		_, _ = fmt.Fscanf(reader, "%d")
		_, _ = fmt.Fscanf(reader, "%d", &dimension)

		maxDim = (1 << dimension) - 1

		fmt.Println("Ingrese un origen: ")
		_, _ = fmt.Fscanf(reader, "%d")
		_, _ = fmt.Fscanf(reader, "%d", &origen)

		if origen > maxDim {
			fmt.Println("El origen es mayor al maximo permitido por la dimension")
			fmt.Println("Presione enter para continuar")
			_, _ = fmt.Fscanf(reader, "%d")
			_, _ = fmt.Fscanf(reader, "%d")
			continue
		}

		fmt.Println("Ingrese un destino: ")
		_, _ = fmt.Fscanf(reader, "%d")
		_, _ = fmt.Fscanf(reader, "%d", &destino)

		if destino > maxDim {
			fmt.Println("El destino es mayor al maximo permitido por la dimension")
			fmt.Println("Presione enter para continuar")
			_, _ = fmt.Fscanf(reader, "%d")
			_, _ = fmt.Fscanf(reader, "%d")
			continue
		}

		camino := pathFinder(int64(origen), int64(destino), dimension)

		for i := range camino {
			node := strconv.FormatInt(camino[i], 2)
			for len(node) < dimension {
				node = "0" + node
			}
			fmt.Printf(node)
			if i < len(camino) - 1 {
				fmt.Print(" -> ")
			} else{
				fmt.Println()
			}
		}
		continuar = false

	}
}

func pathFinder(origen, destino int64, dimension int) []int64 {
	// Declaro la mascara y el slice a devolver
	masks := make([]int64, dimension)
	ret := make([]int64, 0)
	// Agrego el primer nodo que es el origen
	ret = append(ret, origen)
	//Calculo la distancia
	dist := origen ^ destino

	for i := 0; i < dimension; i++ {
		// Variable a devolver
		var node int64
		// Cargo el valor de la mascara en su slice
		masks[i] = 1 << uint8(i)
		// Si tengo que cambiar el bit
		if step := dist & masks[i]; step != 0 {
			//Lo cambio
			if j := len(ret); j == 1 {
				node = origen ^ step
			} else {
				node = ret[j-1] ^ step
			}
			ret = append(ret, node)
		}
	}

	return ret
}

/* ********** */

var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

/*CallClear limpia la consola*/
func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Oops! Su sistema operativo no me dejo limpiar la pantalla")
    }
}