package org.elfwerks.sandbox.javamail;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class ExerciseMail {

	public static final String JAVAMAIL_TEST_CONTEXT = "javamail.xml";
	
	public static void main(String[] args) {

		ApplicationContext context = new ClassPathXmlApplicationContext(JAVAMAIL_TEST_CONTEXT);
		SendMail sendMail = (SendMail)context.getBean("sendMail");
		sendMail.send();
	}

}
