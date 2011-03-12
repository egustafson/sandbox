package org.elfwerks.sandbox.lang.groovy;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

/** Bootstrap class for spring-groovy bean example.
 * @author egustafson
 */
public class Main {
	private static final String appConfig = "applicationContext.xml";

	public static void main(String[] args) {
		System.out.println("Starting Spring-Groovy Bean Example.");
		ApplicationContext ctx = new ClassPathXmlApplicationContext(appConfig);

		
		
		System.out.println("done.");
	}

}
