// PageRequestHandler for Poco WebSockets example
//
/* ====================================================================== */

#include "src/PageRequestHandler.h"

#include <iostream>


using Poco::Net::HTTPRequestHandler;
using Poco::Net::HTTPServerRequest;
using Poco::Net::HTTPServerResponse;


void PageRequestHandler::handleRequest( HTTPServerRequest& request, 
                                        HTTPServerResponse& response ) 
{
    response.setChunkedTransferEncoding(true);
    response.setContentType("text/html");
    std::ostream& ostr = response.send();
    ostr<<"<html>";
    ostr<<"<body>";
    ostr<<"Stub (HTML) output - insert JavaScript HERE";
    ostr<<"</body>";
    ostr<<"</html>";
}
