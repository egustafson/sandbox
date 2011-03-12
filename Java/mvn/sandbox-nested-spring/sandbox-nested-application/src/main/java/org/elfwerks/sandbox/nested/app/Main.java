package org.elfwerks.sandbox.nested.app;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
    private static final Log log = LogFactory.getLog(Main.class);
    private static final String APP_CONTEXT_FILENAME = "applicationContext.xml";
    
    public static void main(String[] args) {
        @SuppressWarnings("unused")
        ApplicationContext cxt = new ClassPathXmlApplicationContext(APP_CONTEXT_FILENAME);
        
        
        log.info("done.");
    }

}
