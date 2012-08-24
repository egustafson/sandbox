/* Template Function
 *
 * This example demonstraits how to make a C++ template with two
 * parameters:  1) Function name, 2) function pointer to the function.
 * The function will take a parameter.
 *
 * The intent here is to create a template class that can be instantiated
 * with a normal ptr to a C function and a name.  Additionally, the
 * instantiated class will be able to register itself in a "table of
 * functions" at program initialization time.  This will allow extension
 * of such a table of functions by linking in additional "function libraries"
 *
 * Author:  Eric Gustafson
 * Date:    3 April 2007
 *
 */
#ifndef _TEMPLATEFUNCTION_HH
#define _TEMPLATEFUNCTION_HH

#include <string>
#include <iostream>

#include "FunctionRegistry.hh"



typedef float (*FnPtr)(float);

class TemplateFunction {
public:
    TemplateFunction(const std::string& name, FnPtr func);
    inline virtual ~TemplateFunction() { };

    float compute(const float val1) { return fnPtr(val1); }
    const std::string& getFnName() { return fnName; }

protected:
    const std::string fnName;
    FnPtr       fnPtr;

private:
    TemplateFunction(const TemplateFunction& right);
    TemplateFunction& operator=(const TemplateFunction& right);
};


TemplateFunction::TemplateFunction(const std::string& name, FnPtr func)
    : fnName(name), fnPtr(func) 
{
    FunctionRegistry::getInstance()->registerFn( this );
    //std::cout<<"Registered: "<<getFnName()<<std::endl;
}

#endif // _TEMPLATEFUNCTION_HH
