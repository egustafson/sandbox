package org.elfwerks.sandbox.spring3.webapp;

import java.util.Map;
import java.util.TreeMap;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RequestParam;

/* This class is Spring session scoped because it holds a bean that is 
 * session scoped.
 */
@Controller
@Scope("session")
@RequestMapping("/list-session-bean")
public class LongListSessionBeanController {
	private static final String controllerPath = "/list-session-bean";

	private int pageSize = 10;
	private int maxIndex = 100;
	
	@Autowired  /* Session scoped in the spring3-servlet.xml bean configuration */
	private PaginationBean paginationBean;
	
	@RequestMapping(method=RequestMethod.GET)
	public String viewPage(@RequestParam(value="next", required=false)String nextPage, ModelMap model) {
		if ( nextPage == null || paginationBean.getPageNumber() > maxIndex ) {
			paginationBean.setPageNumber(0);
		}
		Map<Integer, Integer> data = new TreeMap<Integer, Integer>();
		int nextIndex = paginationBean.getPageNumber() + pageSize;
		for ( int ii = paginationBean.getPageNumber(); ii < nextIndex; ii++ ) {
			data.put(ii, ii+1000);
		}
		model.put("data", data);
		paginationBean.setPageNumber(nextIndex);
		if ( nextIndex < maxIndex ) {
			model.put("morePages", true);
		} 
		model.put("controllerPath", controllerPath);
		return "long-list-session-page";
	}
	
}
