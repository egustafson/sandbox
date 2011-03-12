package org.elfwerks.sandbox.method1;

import org.elfwerks.sandbox.rtcomp.DaemonHolder;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {

	private static final String AppContextFile = "method1/applicationContext.xml";
	
	public static void main(String[] args) {
		ApplicationContext context = new ClassPathXmlApplicationContext(AppContextFile);
		System.out.println("Loaded context: "+context.getDisplayName()+"("+context.getId()+")");
		DaemonHolder dh = context.getBean("daemon", DaemonHolder.class);
		dh.run();
		System.out.println("done.");
	}

}
