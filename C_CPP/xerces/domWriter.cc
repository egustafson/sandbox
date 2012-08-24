#include <xercesc/dom/DOM.hpp>
#include <xercesc/util/XMLString.hpp>
#include <xercesc/util/PlatformUtils.hpp>

#if defined(XERCES_NEW_IOSTREAMS)
#include <iostream>
#else
#include <iostream.h>
#endif

XERCES_CPP_NAMESPACE_USE

int serializeDOM(DOMNode* node) {

    XMLCh tempStr[100];
    XMLString::transcode("LS", tempStr, 99);
    DOMImplementation* impl = DOMImplementationRegistry::getDOMImplementation(tempStr);
    DOMWriter* theSerializer = ((DOMImplementationLS*)impl)->createDOMWriter();

    if (theSerializer->canSetFeature(XMLUni::fgDOMWRTDiscardDefaultContent, true))
        theSerializer->setFeature(XMLUni::fgDOMWRTDiscardDefaultContent, true);

    if (theSerializer->canSetFeature(XMLUni::fgDOMWRTFormatPrettyPrint, true))
        theSerializer->setFeature(XMLUni::fgDOMWRTFormatPrettyPrint, true);

    XMLFormatTarget* myFormTarget = new StdOutFormatTarget();

    try {
        theSeriializer->writeNode(myFormTarget, *node);
    }
    catch (const XMLException& toCatch) {
        char* message = XMLString::transcode(toCatch.getMessage());
        cout << "Exception message is: \n"
             << message << "\n";
        XMLString::release(&message);
        return -1;
    }
    catch (const DOMException& toCatch) {
        char* message = XMLString::transcode(toCatch.msg);
        cout << "Exception message is: \n"
             << message << "\n";
        XMLString::release(&message);
        return -1;
    }
    catch (...) {
        cout << "Unxpected Exception \n";
        return -1;
    }

    theSerializer->release();
    delete myFormTarget;
    return 0;
}

int main() {

    return 0;
}
