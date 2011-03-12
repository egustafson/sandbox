package org.elfwerks.sandbox.utf8;

import java.nio.charset.Charset;

import javax.servlet.http.HttpServletRequest;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;

@Controller
public class FormController {
	private Log log = LogFactory.getLog(this.getClass());

	@RequestMapping(method=RequestMethod.GET, value="/*")
	public String welcome(ModelMap model) {
		log.info("Responding to a GET");
		return "welcome";
	}
	
	@RequestMapping(method=RequestMethod.POST, value="/*")
	public String post(ModelMap model, @RequestParam("value")String value, HttpServletRequest request) {
		log.info("Responding to a POST.");
		log.info("Content-Encoding: "+request.getCharacterEncoding());
		log.info("  value = '"+value+"'");
		StringBuilder sb = new StringBuilder();
		for (int ii = 0; ii < value.length(); ii++) {
			sb.append(value.codePointAt(ii)+" ");
		}
		log.debug("  CodePoints: "+sb.toString());
		byte[] bytes = value.getBytes();
		sb = new StringBuilder();
		for (byte bb : bytes) {
			sb.append(String.format("%02x ", bb));
		}
		log.debug("  Bytes: "+sb.toString());
		String v8 = new String(bytes, Charset.forName("UTF-8"));
		log.info("  v8 = '"+v8+"'");
		sb = new StringBuilder();
		for (int ii = 0; ii < v8.length(); ii++) {
			sb.append(v8.codePointAt(ii)+" ");
		}
		log.debug("  CodePoints: "+sb.toString());
		String characterEncoding = request.getCharacterEncoding();
		
		model.addAttribute("value", value);
		model.addAttribute("tvalue", v8);
		model.addAttribute("encoding", System.getProperty("file.encoding"));
		model.addAttribute("requestEncoding", characterEncoding);
		model.addAttribute("contentType", request.getContentType());
		return "welcome";
	}
}
