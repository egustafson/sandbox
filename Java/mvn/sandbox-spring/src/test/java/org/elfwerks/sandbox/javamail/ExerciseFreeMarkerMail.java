package org.elfwerks.sandbox.javamail;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import org.elfwerks.sandbox.javamail.freemarker.SendmailFreeMarker;

public class ExerciseFreeMarkerMail {

	public static final String JAVAMAIL_TEST_CONTEXT = "javamail.xml";
	
	public static void main(String[] args) {

		ApplicationContext context = new ClassPathXmlApplicationContext(JAVAMAIL_TEST_CONTEXT);
		SendmailFreeMarker sendMail = (SendmailFreeMarker)context.getBean("sendMailFreeMarker");
		sendMail.send();
	}


}
