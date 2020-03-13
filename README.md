# Práctico de máquina

**Ejercicio 1:** Topología Hipercubo

Realice en lenguaje C o Java (usé Golang) un algoritmo capaz de efectuar el routing en una topología de Hipercubo. El algoritmo debe recibir como parámetros la dimensión del hipercubo y los nodos origen y destino. Se debe mostrar por pantalla claramente cuál fue el camino seleccinado (nodos desde el origen hasta el destino). El progrmama deberá ser capaz de realizar el ruteo en cualquier cantidad de dimensiones (parámetro dado).

Para ejecutar el programa en linux se debe correr ```./PM-introSDyP```, en windows ejecutar ```PM-introSDyP.exe```. No es necesario tener golang instalado para ello.

Para correr tanto los test como el benchmark si es necesario tener golang instalado. Para ello ubiquese en .../PM-introSDyP/pathfinder y ejecute el comando ```go test -v --bench . --benchmem```
