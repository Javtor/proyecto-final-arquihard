#!/usr/bin/bash
# Compila el codigo de c++
g++ code/invertir.cpp code/EasyBMP.cpp -o code/invertir
# Ejecuta el coso
code/invertir img/1000.bmp img/1000_inv.bmp 1 1