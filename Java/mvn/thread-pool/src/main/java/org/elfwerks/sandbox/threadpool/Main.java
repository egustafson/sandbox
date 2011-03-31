package org.elfwerks.sandbox.threadpool;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
    private static final Log log = LogFactory.getLog("Application-Main-Class");

    private static final String appConfig = "applicationContext.xml";
    
    public static void main(String[] args) {
        log.info("Start...");
        ApplicationContext ctx = new ClassPathXmlApplicationContext(appConfig);
        log.info("Spring context loaded ["+ctx.getDisplayName()+"].");
        
        PooledWorker worker = ctx.getBean(PooledWorker.class);
        worker.createWork();
        log.info("done.");
    }

}
