/* Function Registry
 *
 * A simple "registry" for TemplateFunction objects.
 *
 * Author:  Eric Gustafson
 * Date:    3 April 2007
 *
 */

#ifndef _FUNCTIONREGISTRY_HH
#define _FUNCTIONREGISTRY_HH

#include <string>
#include <map>

class TemplateFunction;


class FunctionRegistry {
public:
    static FunctionRegistry* getInstance();
    
    TemplateFunction* getFn(const std::string& name);

    void registerFn(TemplateFunction* fn);

protected:
    
    std::map<std::string, TemplateFunction*> registeredFunctions;

private:
    static FunctionRegistry* instance;

    FunctionRegistry() { };
    FunctionRegistry(const FunctionRegistry& right);
    FunctionRegistry& operator=(const FunctionRegistry* right);
    virtual ~FunctionRegistry();
};

#endif // _FUNCTIONREGISTRY_HH
