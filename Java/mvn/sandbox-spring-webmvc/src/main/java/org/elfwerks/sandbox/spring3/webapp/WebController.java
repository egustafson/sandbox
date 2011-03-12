package org.elfwerks.sandbox.spring3.webapp;

import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.stereotype.Controller;

@Controller
//@RequestMapping("/")
public class WebController {

	@ModelAttribute("name")
	public String populateName() {
		return "Name Attribute Value";
	}
	
	@ModelAttribute("refresh")
	public int refreshRate(@RequestParam(value="refresh", defaultValue="0") int refresh) {
		return refresh;
	}
	
	@RequestMapping(method=RequestMethod.GET, value="/*")
	public String welcome(ModelMap model) {
		model.addAttribute("handlerName", "Handler Value");
		return "welcome";
	}
	
	@RequestMapping("/request/{requestId}")
	public String showRequest(@PathVariable("requestId") String requestId, 
			                  ModelMap model) {
		model.addAttribute("requestId", requestId);
		return "show-request";
	}
	
	@RequestMapping(value="/special")
	public String special() {
		return "special";
	}
	
}
