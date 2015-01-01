// Example Poco WebSockets
//
// cloned from the Poco project's WebSocketServer:
// // WebSocketServer.cpp
// //
// // $Id: //poco/1.4/Net/samples/WebSocketServer/src/WebSocketServer.cpp#1 $
// //
// // This sample demonstrates the WebSocket class.
// //
// // Copyright (c) 2012, Applied Informatics Software Engineering GmbH.
// // and Contributors.
// //
// // SPDX-License-Identifier: BSL-1.0
// //
/* ====================================================================== */

#include "Poco/Net/HTTPServer.h"
#include "Poco/Net/HTTPRequestHandler.h"
#include "Poco/Net/HTTPRequestHandlerFactory.h"
#include "Poco/Net/HTTPServerParams.h"
#include "Poco/Net/HTTPServerRequest.h"
#include "Poco/Net/HTTPServerResponse.h"
#include "Poco/Net/HTTPServerParams.h"
#include "Poco/Net/ServerSocket.h"
#include "Poco/Net/WebSocket.h"
#include "Poco/Net/NetException.h"
#include "Poco/Util/ServerApplication.h"
#include "Poco/Util/Option.h"
#include "Poco/Util/OptionSet.h"
#include "Poco/Util/HelpFormatter.h"
#include "Poco/Format.h"
#include <iostream>

using Poco::Net::ServerSocket;
using Poco::Net::WebSocket;
using Poco::Net::WebSocketException;
using Poco::Net::HTTPRequestHandler;
using Poco::Net::HTTPRequestHandlerFactory;
using Poco::Net::HTTPServer;
using Poco::Net::HTTPServerRequest;
using Poco::Net::HTTPResponse;
using Poco::Net::HTTPServerResponse;
using Poco::Net::HTTPServerParams;
using Poco::Timestamp;
using Poco::ThreadPool;
using Poco::Util::ServerApplication;
using Poco::Util::Application;
using Poco::Util::Option;
using Poco::Util::OptionSet;
using Poco::Util::HelpFormatter;

#include "config.h"  // from autoconf


// PageRequestHandler - return the web page.
//
class PageRequestHandler: public HTTPRequestHandler {
public:
    void handleRequest(HTTPServerRequest& request, HTTPServerResponse& response) {
        response.setChunkedTransferEncoding(true);
        response.setContentType("text/html");
        std::ostream& ostr = response.send();
        ostr<<"<html>";
        ostr<<"<body>";
        ostr<<"Stub (HTML) output - insert JavaScript HERE";
        ostr<<"</body>";
        ostr<<"</html>";
    }
}; // end class PageRequestHandler


// WebSocketRequestHandler - handle the WebSocket connection.
//
class WebSocketRequestHandler: public HTTPRequestHandler {
public:
    void handleRequest(HTTPServerRequest& request, HTTPServerResponse& response) {
        Application& app = Application::instance();
        //
        // TBD
        //
        app.logger().information("Stub - handling WebSocket request");
        response.setStatusAndReason(HTTPResponse::HTTP_BAD_REQUEST);
        response.setContentLength(0);
        response.send();
    }
}; // end class WebSocketRequestHandler


// RequestHandlerFactory - produce ReqHdlr object based on URI.
//
class RequestHandlerFactory: public HTTPRequestHandlerFactory {
public:
    HTTPRequestHandler* createRequestHandler(const HTTPServerRequest& request) {
        Application& app = Application::instance();
        app.logger().information("Request from "
                                 + request.clientAddress().toString() + ": " 
                                 + request.getMethod() + " "
                                 + request.getURI() + " "
                                 + request.getVersion());
        for (HTTPServerRequest::ConstIterator it = request.begin(); it != request.end(); ++it) {
            app.logger().information(it->first + ": " + it->second);
        }
        app.logger().information("---");
        if ( request.find("Upgrade") != request.end() && Poco::icompare(request["Upgrade"], "websocket") == 0) {
            return new WebSocketRequestHandler;
        } else {
            return new PageRequestHandler;
        }
    }
}; // end class RequestHandlerFactory


//
// WebSocketServer (main)
//
class WebSocketServer : public Poco::Util::ServerApplication {
public:
    WebSocketServer() { /* nothing to do */ }
    ~WebSocketServer() { /* nothing to do */ }

protected:
    void initialize(Application& self) {
        // loadConfiguration():
        ServerApplication::initialize(self);
    }

    void uninitialze() {
        ServerApplication::uninitialize();
    }

    int main(const std::vector<std::string>& args) {
        std::cout<<"Package '"<<PACKAGE_STRING<<"'"<<std::endl;

        unsigned short port = 9980;
        ServerSocket svs(port);
        HTTPServer srv(new RequestHandlerFactory, svs, new HTTPServerParams);
        srv.start();
        std::cout<<"serving..."<<std::endl;
        waitForTerminationRequest();  // ctrl-c or SIGTERM
        srv.stop();

        std::cout<<"done."<<std::endl;
        return Application::EXIT_OK;
    }

}; // end class WebSocketServer

POCO_SERVER_MAIN(WebSocketServer)

// Local Variables:
// mode: C++
// End:
