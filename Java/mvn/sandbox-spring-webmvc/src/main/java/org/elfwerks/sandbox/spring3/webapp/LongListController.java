package org.elfwerks.sandbox.spring3.webapp;

import java.util.Map;
import java.util.TreeMap;

import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
@RequestMapping("/list")
public class LongListController {

	int pageSize = 10;
	
	@RequestMapping(method=RequestMethod.GET)
	public String initialPage(ModelMap model) {
		return viewPage(0, model);
	}
	
	@RequestMapping(method=RequestMethod.GET, value="/{index}")
	public String viewPage(@PathVariable("index") Integer index, ModelMap model) {
		Map<Integer, Integer> data = new TreeMap<Integer, Integer>();
		Integer nextPageStartIndex = index + pageSize;
		for (int ii = index; ii < nextPageStartIndex; ii++) {
			data.put(ii, ii+1000);
		}
		if ( nextPageStartIndex > 100 ) {
			nextPageStartIndex = null;
		}
		model.put("data", data);
		model.put("nextIndex", nextPageStartIndex);
		return "long-list-page";
	}
	
}
