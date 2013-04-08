package org.elfwerks.sandbox.spring.property;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.context.support.AbstractApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
	static final String contextFile = "application-context.xml";
	static final Log log = LogFactory.getLog("Main");


	public static void main(String[] args) {
		AbstractApplicationContext context = new ClassPathXmlApplicationContext(contextFile);
		InjectionBean ib = context.getBean("ibean", InjectionBean.class);
		log.info(ib.toString());
		context.close();
		log.info("done.");

	}

}
