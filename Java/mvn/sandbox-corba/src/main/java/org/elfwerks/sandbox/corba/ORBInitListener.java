package org.elfwerks.sandbox.corba;

import java.util.Properties;
import javax.servlet.ServletContextEvent;
import javax.servlet.ServletContextListener;
import org.omg.CORBA.ORB;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class ORBInitListener implements ServletContextListener {
    private static final Log log = LogFactory.getLog(ORBInitListener.class);
    
    private class OrbRunner extends Thread {
        ORB theOrb;
        public OrbRunner(ORB orb) {
            super("orb-runner");
            theOrb = orb;
        }
        public void run() {
            theOrb.run();
        }
    }

    @Override
    public void contextInitialized(ServletContextEvent sce) {
        log.info("Starting the ORB.");
        Properties orbProperties = new Properties();
        orbProperties.setProperty("org.omg.CORBA.ORBClass", "org.jacorb.orb.ORB");
        orbProperties.setProperty("org.omg.CORBA.ORBSingletonClass", "org.jacorb.orb.ORBSingleton");
        String [] args = new String[0];
        ORB orb = ORB.init(args, orbProperties);
        sce.getServletContext().setAttribute("ORB", orb);
        log.info("ORB initialized.");
        OrbRunner orbRunner = new OrbRunner(orb);
        orbRunner.start();
        sce.getServletContext().setAttribute("ORBThread", orbRunner);
        log.info("ORB running.");
    }

    @Override
    public void contextDestroyed(ServletContextEvent sce) {
        log.info("Shutting down the ORB.");
        ORB orb = (ORB)sce.getServletContext().getAttribute("ORB");
        orb.shutdown(true);
        OrbRunner orbRunner = (OrbRunner)sce.getServletContext().getAttribute("ORBThread");
        try {
            orbRunner.join();
        } catch (InterruptedException ex) {
            log.warn("Interrupted while attempting to join() with the orb.run() thread.");
        }
        log.info("ORB shutdown complete.");
    }

}
