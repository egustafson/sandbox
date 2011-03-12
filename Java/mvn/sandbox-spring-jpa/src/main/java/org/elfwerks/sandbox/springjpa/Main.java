package org.elfwerks.sandbox.springjpa;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.context.support.AbstractApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

/** Demo a J2SE Spring enabled JPA Initialization
 */
public class Main {
	private static String ContextFileName = "META-INF/applicationContext.xml";
	private static Log log = LogFactory.getLog(Main.class);
	
    private static AbstractApplicationContext applicationContext;

	
    public static void main( String[] args ) {
    	applicationContext = initializeContext();
    	try {
    		TxWorkerBean worker = (TxWorkerBean)applicationContext.getBean("exampleTxBean");
    		worker.lookupNoTx();
    		worker.modifyRequireTx();
    		worker.nestedRequireTx();
    	} finally {
    		applicationContext.close();
    	}
        log.info("Done.");
    }
    
    
    
    /** Initialize the Spring <code>ApplicationContext</code>.  
     * This method is public only to allow using it during JUnit testing.
     */
    public static AbstractApplicationContext initializeContext() {
        AbstractApplicationContext context = new ClassPathXmlApplicationContext(ContextFileName);
        //log.info("Application Context initialized.");
        return context;
    }

}
