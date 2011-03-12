package org.elfwerks.sandbox.javamail;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.mail.MailException;
import org.springframework.mail.MailSender;
import org.springframework.mail.SimpleMailMessage;

public class SendMail {
	private static final Log log = LogFactory.getLog(SendMail.class);
	
	private MailSender mailSender;
	private SimpleMailMessage templateMessage;
	
	public void setMailSender(MailSender sender) {
		mailSender = sender;
	}
	
	public void setTemplateMessage(SimpleMailMessage template) {
		templateMessage = template;
	}
	
	public void send() {
		
		SimpleMailMessage msg = new SimpleMailMessage(templateMessage);
		msg.setTo("ericg@elfwerks.org");
		msg.setText("Test message.");
		try {
			mailSender.send(msg);
		} catch (MailException ex) {
			log.error("Failed to send message - MailException.  (reason:"+ex.getLocalizedMessage()+")");
		}
		
	}
	
}
