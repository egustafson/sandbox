
#include <Python.h>

#include <stdio.h>

#define SCRIPT_NAME "demo.py"

int main(int argc, char* argv[]) {

    FILE* fp;

    fp = fopen(SCRIPT_NAME, "r");

    Py_SetProgramName(argv[0]);
    Py_Initialize();

    PyRun_SimpleFile(fp, SCRIPT_NAME);

    Py_Finalize();
    return 0;
}
