package main

import (
	"bufio"
	"fmt"
	pf "github.com/FacundoDecena/PM-introSDyP/pathfinder"
	"os"
	"os/exec"
	"runtime"
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

		fmt.Println(pf.PathFinder(int64(origen), int64(destino), dimension))

		continuar = false

	}
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
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Oops! Su sistema operativo no me dejo limpiar la pantalla")
	}
}
