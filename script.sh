#!/usr/bin/bash
# Compila el codigo de c++
go get "golang.org/x/image/bmp"
go build src/go/invertir_go.go
g++ src/cpp/invertir_cpp.cpp src/cpp/EasyBMP.cpp -o invertir_cpp
# Ejecuta c++
./invertir_cpp img/400.bmp img/400_inv.bmp 1 1
./invertir_go 2 2 1 500