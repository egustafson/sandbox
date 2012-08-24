#include <stdio.h>
#include <xercesc/util/PlatformUtils.hpp>

XERCES_CPP_NAMESPACE_USE

int main(int argc, char* argv[])
{
    try {
        XMLPlatformUtils::Initialize();
    }
    catch ( const XMLException& toCatch ) {
        // Failure Processing.
        fprintf(stderr, "Caught an XMLException\n");
        return 1;
    }
    
    XMLPlatformUtils::Terminate();

    return 0;
}
