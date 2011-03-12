package org.elfwerks.sandbox.corba;

import java.io.IOException;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.omg.CORBA.Object;
import org.omg.CORBA.ORB;
import org.omg.PortableServer.POA;
import org.omg.PortableServer.POAHelper;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

/**
 * Servlet implementation class for Servlet: PingServantServlet
 *
 */
public class PingServantServlet extends javax.servlet.http.HttpServlet implements javax.servlet.Servlet {
    private static final Log log = LogFactory.getLog(PingServantServlet.class);
    static final long serialVersionUID = 1L;
    
    String serviceIor;

    /**
     * @see javax.servlet.GenericServlet#init()
     */
    @Override
    public void init() throws ServletException {
        try {
            ORB orb = (ORB)getServletContext().getAttribute("ORB");
            POA poa = POAHelper.narrow(orb.resolve_initial_references("RootPOA"));
            poa.the_POAManager().activate();
            Object pingRef = poa.servant_to_reference(new PingServant());
            serviceIor = orb.object_to_string(pingRef);
            log.info("CORBA PingService registered.");
        } catch (Exception ex) {
            String msg = "The CORBA service 'PingService' failed to initialize.";
            log.fatal(msg);
            throw new ServletException(msg, ex);
        }
    }   

    /**
     * @see javax.servlet.http.HttpServlet#doGet(HttpServletRequest request, HttpServletResponse response)
     */
    @Override
    protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        response.setContentType("text/plain");
        response.getWriter().print(serviceIor);
    }

    /**
     * @see javax.servlet.GenericServlet#destroy()
     */
    @Override
    public void destroy() {
        log.info("shutdown.");
    }

}