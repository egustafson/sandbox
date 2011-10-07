package org.elfwerks.demo.webapp;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.servlet.ModelAndView;

@Controller
@RequestMapping("/hello")
public class ExampleController {
	private final Log log = LogFactory.getLog(this.getClass());
	
	private String message = "Default Message";
	public void setMessage(String m) { message = m; }
	
	@RequestMapping(method=RequestMethod.GET)
	public String get() {
		log.info("processing GET request for /hello");
			return "hello-view";
	}
	
//	@RequestMapping(method=RequestMethod.GET)
//	public ModelAndView get() {
//		log.info("processing GET request for /hello");
//		ModelAndView mv = new ModelAndView("hello-view");
//		mv.addObject("msg", message);
//		return mv;
//	}
	
}
