package org.elfwerks.sandbox.spring3.webapp;

import java.util.Map;
import java.util.TreeMap;

import javax.servlet.http.HttpSession;

import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
@RequestMapping("/list-session-attr")
public class LongListSessionAttrController {
	private static final String pbname = "paginationBean";
	private static final String controllerPath = "/list-session-attr";

	private int pageSize = 10;
	private int maxIndex = 100;
	
	@RequestMapping(method=RequestMethod.GET)
	public String viewPage(HttpSession session, ModelMap model) {
		PaginationBean page = (PaginationBean)session.getAttribute(pbname);
		if ( page == null ) {
			page = new PaginationBean();
			session.setAttribute(pbname, page);
		}
		Map<Integer, Integer> data = new TreeMap<Integer, Integer>();
		int nextIndex = page.getPageNumber() + pageSize;
		for ( int ii = page.getPageNumber(); ii < nextIndex; ii++ ) {
			data.put(ii, ii+1000);
		}
		model.put("data", data);
		page.setPageNumber(nextIndex);
		model.put("morePages", true);
		model.put("controllerPath", controllerPath);
		if ( nextIndex >= maxIndex ) {
			session.removeAttribute(pbname);
		}
		return "long-list-session-page";
	}
	
}
