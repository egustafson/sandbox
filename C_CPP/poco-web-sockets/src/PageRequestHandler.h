// PageRequestHandler for Poco WebSockets example
//
// cloned from the Poco project's WebSocketServer
// Modified to read "pages" from file.
//
/* ====================================================================== */

#include "Poco/Net/HTTPRequestHandler.h"
#include "Poco/Net/HTTPServerRequest.h"
#include "Poco/Net/HTTPServerResponse.h"



class PageRequestHandler: public Poco::Net::HTTPRequestHandler {
 public:
    void handleRequest( Poco::Net::HTTPServerRequest& request, 
                        Poco::Net::HTTPServerResponse& response );
};


// Local Variables:
// mode: C++
// End:
