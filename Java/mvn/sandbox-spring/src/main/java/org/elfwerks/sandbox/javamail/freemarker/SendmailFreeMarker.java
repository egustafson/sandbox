package org.elfwerks.sandbox.javamail.freemarker;

import java.io.StringWriter;
import java.io.Writer;
import java.util.HashMap;
import java.util.Map;
import javax.mail.internet.MimeMessage;

import freemarker.template.Configuration;
import freemarker.template.Template;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.mail.javamail.MimeMessageHelper;
import org.springframework.mail.javamail.MimeMessagePreparator;


public class SendmailFreeMarker {
	private JavaMailSender mailSender;
	private Configuration freemarkerConfig;
	
	public void setMailSender(JavaMailSender mailSender) {
		this.mailSender = mailSender;
	}
	
	public void setFreemarkerConfig(Configuration freemarkerConfig) {
		this.freemarkerConfig = freemarkerConfig;
	}

	public void send() {
		MimeMessagePreparator preparator = new MimeMessagePreparator() {
			public void prepare(MimeMessage mimeMessage) throws Exception {
				MimeMessageHelper message = new MimeMessageHelper(mimeMessage);
				message.setTo("ericg@elfwerks.org");
				message.setFrom("ericg@elfwerks.org");
				message.setSubject("FreeMarker HTML Demo");
				Map<String, Object> model = new HashMap<String, Object>();
				model.put("greeting", "Hello from the java program, this is a FreeMarker demo.");
				
				Template tmpl = freemarkerConfig.getTemplate("mail-tmpl.ftl");
				Writer out = new StringWriter();
				tmpl.process(model, out);
				message.setText(out.toString(), true);
			}
		};
		this.mailSender.send(preparator);
	}

}
