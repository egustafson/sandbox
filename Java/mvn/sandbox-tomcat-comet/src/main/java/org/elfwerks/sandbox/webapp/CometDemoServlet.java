package org.elfwerks.sandbox.webapp;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;

import org.apache.catalina.CometEvent;
import org.apache.catalina.CometProcessor;

public class CometDemoServlet extends HttpServlet implements CometProcessor {
    private static final long serialVersionUID = 8000274619623300354L;

    
    
    @Override
    public void event(CometEvent event) throws IOException, ServletException {
        System.out.println("Received COMET Event: "+event.getEventType());
        event.close();
    }

}
