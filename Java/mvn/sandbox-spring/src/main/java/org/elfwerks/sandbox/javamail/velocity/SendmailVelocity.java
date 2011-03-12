package org.elfwerks.sandbox.javamail.velocity;

import javax.mail.internet.MimeMessage;
import java.util.HashMap;
import java.util.Map;

import org.apache.velocity.app.VelocityEngine;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.mail.javamail.MimeMessagePreparator;
import org.springframework.ui.velocity.VelocityEngineUtils;

public class SendmailVelocity {

	private JavaMailSender mailSender;
	private VelocityEngine velocityEngine;
	
	public void setMailSender(JavaMailSender mailSender) {
		this.mailSender = mailSender;
	}
	
	public void setVelocityEngine(VelocityEngine velocityEngine) {
		this.velocityEngine = velocityEngine;
	}

	public void send() {
		MimeMessagePreparator preparator = new MimeMessagePreparator() {
			public void prepare(MimeMessage mimeMessage) throws Exception {
				MimeMessageHelper message = new MimeMessageHelper(mimeMessage);
				message.setTo("ericg@elfwerks.org");
				message.setFrom("ericg@elfwerks.org");
				message.setSubject("Velocity HTML Demo");
				Map<String, Object> model = new HashMap<String, Object>();
				model.put("greeting", "Hello from the java program, this is a Velocity demo.");
				String text = VelocityEngineUtils.mergeTemplateIntoString(velocityEngine, "velocity/mail-tmpl.vm", model);
				message.setText(text, true);
			}
		};
		this.mailSender.send(preparator);
	}
	
}
