package org.elfwerks.sandbox.spring3.webapp;

import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;

@Controller
@RequestMapping("/alt")
public class AltController {


	@RequestMapping(method=RequestMethod.GET)
	public String getRequest(ModelMap model) {
		return "alt";
	}
	
	@RequestMapping(method=RequestMethod.GET, params="p")
	public String paramRequest(@RequestParam("p") String p, ModelMap model) {
		String message = "The parameter value was [" + p + "]";
		model.addAttribute("message", message);
		return "alt";
	}
	
}
