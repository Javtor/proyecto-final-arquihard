#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main()
{
    freopen("script.sh", "w", stdout);
    cout << "#!/usr/bin/bash\n"
         << "# Compila el codigo de c++\n"
         << "go get \"golang.org/x/image/bmp\"\n"
         << "go build src/go/invertir_go.go\n"
         << "g++ src/cpp/invertir_cpp.cpp src/cpp/EasyBMP.cpp -o invertir_cpp\n\n";

    vector<string> lenguajes = {"invertir_cpp", "invertir_go"};
    vector<string> imgs = {"400", "500", "600", "700", "800", "900", "1000", "1100"};

    int pc = 1;
    int tratamientos = 3;
    int algoritmos = 5;

    vector<string> comandos;
    for (auto lang : lenguajes)
    {
        for (int alg = 1; alg <= algoritmos; alg++)
        {
            for (int trat = 1; trat <= tratamientos; trat++)
            {
                for (auto img : imgs)
                {
                    string comando = "./" + lang + "  " + to_string(pc) + " " + to_string(alg) + " " + to_string(trat) + " " + img;
                    comandos.push_back(comando);
                }
            }
        }
    }
    random_shuffle(comandos.begin(), comandos.end());
    for (auto com : comandos)
    {
        cout << com << endl;
    }
    return 0;
}
