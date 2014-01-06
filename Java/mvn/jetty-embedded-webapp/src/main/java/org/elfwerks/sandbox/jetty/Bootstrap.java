package org.elfwerks.sandbox.jetty;

import java.lang.management.ManagementFactory;
import java.net.URL;
import java.security.ProtectionDomain;

import org.eclipse.jetty.jmx.MBeanContainer;
import org.eclipse.jetty.server.Server;
import org.eclipse.jetty.webapp.WebAppContext;

public class Bootstrap {

	public static void main(String[] args) throws Exception {
	    Server server = new Server(8080);

	    ProtectionDomain domain = Bootstrap.class.getProtectionDomain();
	    URL myLocation = domain.getCodeSource().getLocation();
	    
	    // Setup WebAppContext
	    WebAppContext webapp = new WebAppContext();
	    webapp.setContextPath("/");
	    webapp.setDescriptor(myLocation.toExternalForm() + "/WEB-INF/web.xml");
	    webapp.setServer(server);
	    webapp.setWar(myLocation.toExternalForm());
	    
	    server.setHandler(webapp);
	    
        // Setup JMX
        MBeanContainer mbContainer = new MBeanContainer(ManagementFactory.getPlatformMBeanServer());
        server.addBean(mbContainer);
        
	    server.start();
	    server.join();
	}

}
