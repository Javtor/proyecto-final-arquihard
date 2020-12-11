# Para ejecutar la versión de c++
Usando g++ en ubuntu

Compilar: `g++ code/invertir.cpp code/EasyBMP.cpp -o code/invertir`

Ejecutar: `./invertir ruta_origen ruta_destino version tratamiento`

Ejemplo: `src/cpp/invertir img/1000.bmp img/1000_inv.bmp 1 1`



Usando go en linux

se le pasan argumentos ./invert pc versionAlgoritmo tratamiento versionImg

Antes de : `go get "golang.org/x/image/bmp"`

Compilar: `go build invertir.go`

Ejecutar: `./invert pc versionAlgoritmo tratamiento versionImg`

Ejemplo `src/go/invertir_go 2 2 1 400`
