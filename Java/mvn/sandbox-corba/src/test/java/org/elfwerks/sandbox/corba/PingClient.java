package org.elfwerks.sandbox.corba;

import java.net.URL;
import java.net.URLConnection;
import java.util.Properties;
import org.omg.CORBA.ORB;
import org.omg.CORBA.Object;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import org.elfwerks.sandbox.servant.PingService;
import org.elfwerks.sandbox.servant.PingServiceHelper;

public class PingClient {
    private static final Log log = LogFactory.getLog(PingClient.class);
    private static final String pingIor = "http://localhost:8080/sandbox-corba/ping-ior";

    public static void main(String args[]) {
        log.info("Starting.");
        try {
            ORB orb = initOrb();
            String ior = getIor();
            Object pingRef = orb.string_to_object(ior);
            PingService ping = PingServiceHelper.narrow(pingRef);
            
            ping.ping();
            log.info("Ping successful.");
        } catch (Exception ex) {
            log.fatal("client failed.", ex);
        }
    }

    private static String getIor() {
        try {
            URLConnection connection = (new URL(pingIor)).openConnection();
            int contentLength = connection.getContentLength();
            byte[] content = new byte[contentLength];
            connection.getInputStream().read(content);
            String ior = new String(content);
            return ior;
        } catch (Exception ex) {
            log.fatal(ex);
        }
        return null;
    }
    
    private static ORB initOrb() {
        Properties orbProperties = new Properties();
        orbProperties.setProperty("org.omg.CORBA.ORBClass", "org.jacorb.orb.ORB");
        orbProperties.setProperty("org.omg.CORBA.ORBSingletonClass", "org.jacorb.orb.ORBSingleton");
        String [] args = new String[0];
        ORB orb = ORB.init(args, orbProperties);
        return orb;
    }
    
}
