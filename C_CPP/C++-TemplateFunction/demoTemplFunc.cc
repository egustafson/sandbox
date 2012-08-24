/* Template Function - Main Program
 *
 * See TemplateFunction.hh for details.
 *
 * Author:  Eric Gustafson
 * Date:    3 April 2007
 *
 */
#include <string>
#include <iostream>
#include "TemplateFunction.hh"

typedef float (*FpPtr)(float);

float demoFunc(float val) {
    return val/2.0;
}

TemplateFunction fn("demo-func", demoFunc);

int main() {

    FpPtr func = demoFunc;
    float x = (*func)(3);
    printf("x = %f\n", x);

    std::cout<<fn.compute(4.0)<<std::endl;

    return 0;
}
