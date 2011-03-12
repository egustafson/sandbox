package org.elfwerks.sandbox.spring3.webapp;

import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;

@Controller
@RequestMapping("/search-demo")
public class SearchController {

	@RequestMapping(method=RequestMethod.GET)
	public String get() {
		return "search-page";
	}
	
	@RequestMapping(method=RequestMethod.POST, params="searchKey")
	public String processForm(@RequestParam("searchKey") String searchKey) {
		return "redirect:/search-demo/"+searchKey;
	}
	
	@RequestMapping(value="/{key}", method=RequestMethod.GET)
	public String search(@PathVariable("key") String key, ModelMap model) {
		model.addAttribute("searchKey", key);
		model.addAttribute("searchResult", "You searched for '"+key+"'");
		return "search-page";
	}
	
}
