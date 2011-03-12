package org.elfwerks.sandbox.pool;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
	private static final String appConfig = "applicationContext.xml";

	public static void main(String[] args) {
		ApplicationContext ctx = new ClassPathXmlApplicationContext(appConfig);
		System.out.println("Loaded context: "+ctx.getDisplayName()+"("+ctx.getId()+")");

		
		
		PoolManager mgr = ctx.getBean(PoolManager.class);
		mgr.doProcessing();
		System.out.println("done.");
	}

}
