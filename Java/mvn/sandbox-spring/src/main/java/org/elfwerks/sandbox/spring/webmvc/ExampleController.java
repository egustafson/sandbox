package org.elfwerks.sandbox.spring.webmvc;

import java.util.Date;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.mvc.Controller;

public class ExampleController implements Controller {

	@Override
	public ModelAndView handleRequest(HttpServletRequest request, HttpServletResponse response) throws Exception {
		List<String> messages = new LinkedList<String>();
		messages.add("Message 1");
		messages.add("Message two");
		messages.add("Message number three");
		messages.add("Final Message");
		
		Map<String, Object> model = new HashMap<String, Object>();
		model.put("hello", "Hello from the Controller");
		model.put("messages", messages);
		model.put("datetime", new Date());
		ModelAndView mv = new ModelAndView("demoview", model);
		return mv;
	}

}
