package org.elfwerks.sandbox.javamail;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import org.elfwerks.sandbox.javamail.velocity.SendmailVelocity;

public class ExerciseVelocityMail {

	public static final String JAVAMAIL_TEST_CONTEXT = "javamail.xml";
	
	public static void main(String[] args) {

		ApplicationContext context = new ClassPathXmlApplicationContext(JAVAMAIL_TEST_CONTEXT);
		SendmailVelocity sendMail = (SendmailVelocity)context.getBean("sendMailVelocity");
		sendMail.send();
	}

}
