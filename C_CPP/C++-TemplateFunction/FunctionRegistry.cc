/* Function Registry
 *
 * Author:  Eric Gustafson
 * Date:    3 April 2007
 */

#include "FunctionRegistry.hh"
#include "TemplateFunction.hh"

#include <iostream>

FunctionRegistry* FunctionRegistry::getInstance() {
    if ( !instance ) {
        instance = new FunctionRegistry();
    }
    return instance;
}


TemplateFunction* FunctionRegistry::getFn(const std::string& name) {
    TemplateFunction* tf = NULL;
    std::map<std::string, TemplateFunction*>::iterator it = registeredFunctions.find(name);
    if ( it != registeredFunctions.end() ) {
        tf = it->second;
    }
    return tf;
}

void FunctionRegistry::registerFn(TemplateFunction* fn) {
    registeredFunctions[fn->getFnName()] = fn;
    std::cout<<"Registered: "<<fn->getFnName()<<std::endl;
}
