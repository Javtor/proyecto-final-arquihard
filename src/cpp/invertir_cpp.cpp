#pragma GCC optimize("O2") 
#include <iostream>
#include <stdlib.h>
#include "EasyBMP.h"
#include <cstring>
#include <chrono>
#include <fstream>
#include <vector>
#include <numeric>

using namespace std;

struct rgb
{
    uint8_t r, g, b;
};

const int NUM_MUESTRAS = 100;
const double confianza = 1.96; //95%
const double error = 0.25;

int main(int argc, char *argv[])
{

    if (argc < 5)
    {
        std::cout << "Verifica los argumentos de consola\nPresiona ENTER para continuar...\n";
        getchar();
        return 0;
    }

    auto argpc = argv[1];
    auto argversion = argv[2];
    auto argtratamiento = argv[3];
    auto argorigen = argv[4];

    fstream apilados;
    apilados.open("apilados.csv", fstream::app);

    fstream metricas;
    metricas.open("metricas.csv", fstream::app);

    if(!(apilados.good())){
        apilados << "PC;Tam (nxn);Algoritmo;Lenguaje;Repeticion;Tiempo (ns)";
    }

    // Lee el bmp dentro de la matriz de pixeles
    BMP img;
    string srcPath = "img/" + string(argorigen) + ".bmp";
    string dstPath = "img/inverted_" + string(argorigen) + ".bmp";
    img.ReadFromFile(srcPath.c_str());
    const int width = img.TellWidth();
    const int height = img.TellHeight();
    rgb ImRGB0[height][width];

    for (int r = 0; r < height; r++)
    {
        for (int c = 0; c < width; c++)
        {
            RGBApixel Temp = img.GetPixel(c, r);
            ImRGB0[r][c].r = (int)Temp.Red;
            ImRGB0[r][c].g = (int)Temp.Green;
            ImRGB0[r][c].b = (int)Temp.Blue;
        }
    }

    rgb ImRGB[height][width];
    memcpy(ImRGB, ImRGB0, height * width * sizeof(rgb));

    //Ejecuta el algoritmo
    long version = strtol(argversion, NULL, 10);
    string tratamiento(argtratamiento);
    string pc(argpc);
    string archivoSalida = "data/pc" + pc + "-cpp-" + to_string(width) + "-version" + to_string(version) + "-tratamiento" + tratamiento + ".txt";
    freopen(archivoSalida.c_str(), "w", stdout);
    vector<double> datos;
    int n = NUM_MUESTRAS;
    while (n--)
    {
        auto start = std::chrono::high_resolution_clock::now();
        switch (version)
        {
        case 1:
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].r = 255 - ImRGB[r][c].r;
                    ImRGB0[r][c].g = 255 - ImRGB[r][c].g;
                    ImRGB0[r][c].b = 255 - ImRGB[r][c].b;
                }
            }
            break;

        case 2:
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].r = 255 - ImRGB[r][c].r;
                }
            }
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].g = 255 - ImRGB[r][c].g;
                }
            }
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].b = 255 - ImRGB[r][c].b;
                }
            }
            break;

        case 3:
            for (int c = 0; c < width; c++)
            {
                for (int r = 0; r < height; r++)
                {
                    ImRGB0[r][c].r = 255 - ImRGB[r][c].r;
                    ImRGB0[r][c].g = 255 - ImRGB[r][c].g;
                    ImRGB0[r][c].b = 255 - ImRGB[r][c].b;
                }
            }
            break;
        case 4:
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].r = 255 - ImRGB[r][c].r;
                }
            }
            for (int r = 0; r < height; r++)
            {
                for (int c = 0; c < width; c++)
                {
                    ImRGB0[r][c].g = 255 - ImRGB[r][c].g;
                    ImRGB0[r][c].b = 255 - ImRGB[r][c].b;
                }
            }
            break;
        case 5:
            for (int r = 0; r < height; r += 2)
            {
                for (int c = 0; c < width; c += 2)
                {
                    ImRGB0[r][c].r = 255 - ImRGB[r][c].r;
                    ImRGB0[r][c].g = 255 - ImRGB[r][c].g;
                    ImRGB0[r][c].b = 255 - ImRGB[r][c].b;

                    ImRGB0[r][c + 1].r = 255 - ImRGB[r][c + 1].r;
                    ImRGB0[r][c + 1].g = 255 - ImRGB[r][c + 1].g;
                    ImRGB0[r][c + 1].b = 255 - ImRGB[r][c + 1].b;

                    ImRGB0[r + 1][c].r = 255 - ImRGB[r + 1][c].r;
                    ImRGB0[r + 1][c].g = 255 - ImRGB[r + 1][c].g;
                    ImRGB0[r + 1][c].b = 255 - ImRGB[r + 1][c].b;

                    ImRGB0[r + 1][c + 1].r = 255 - ImRGB[r + 1][c + 1].r;
                    ImRGB0[r + 1][c + 1].g = 255 - ImRGB[r + 1][c + 1].g;
                    ImRGB0[r + 1][c + 1].b = 255 - ImRGB[r + 1][c + 1].b;
                }
            }
            break;

        default:
            break;
        }
        auto stop = std::chrono::high_resolution_clock::now();
        auto duration = std::chrono::duration_cast<std::chrono::nanoseconds>(stop - start);
        auto cnt = (double)duration.count();
        double normalized = cnt / (double)(width * height);
        std::cout << normalized << std::endl;
        datos.push_back(normalized);
        apilados << pc + ";" + argorigen + ";" + to_string(version) + ";c++;" + argtratamiento + ";" + to_string(normalized) << std::endl;
    }

    double sum = 0.0;
    for(double dato : datos){
        sum += dato;
    }
    double media = sum/datos.size();
    double varianza = 0;
    for(double dato : datos){
        varianza += (dato-media)*(dato-media);
    }
    varianza /= (datos.size()-1);
    double desv = sqrt(varianza);
    double tamMuestra = (confianza*desv/error)*(confianza*desv/error);


     metricas<< pc + ";" + argorigen + ";" + to_string(version) + ";c++;" + argtratamiento +';'+ to_string(media) +';'+to_string(varianza) +';'+to_string(desv) +';'+to_string(tamMuestra)<< std::endl;

    //Escribe la matriz de pixeles en el nuevo bmp
    for (int r = 0; r < height; r++)
    {
        for (int c = 0; c < width; c++)
        {
            RGBApixel Temp;
            Temp.Red = ImRGB0[r][c].r;
            Temp.Green = ImRGB0[r][c].g;
            Temp.Blue = ImRGB0[r][c].b;
            img.SetPixel(c, r, Temp);
        }
    }
    img.WriteToFile(dstPath.c_str());

    return 0;
}
