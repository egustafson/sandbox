package org.elfwerks.sandbox.quartz.webapp;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.quartz.Scheduler;
import org.quartz.SchedulerException;
import org.quartz.SchedulerFactory;
import org.quartz.impl.DirectSchedulerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.ModelMap;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class ApplicationController {
	private final Log log = LogFactory.getLog(this.getClass());

	@Autowired
	private Scheduler scheduler;
	
	@RequestMapping(method=RequestMethod.GET, value="/*")
	public String welcome(ModelMap model) {
		model.addAttribute("scheduler", scheduler);
		try {
			log.info("Scheduler Context: "+scheduler.getContext().toString());
			for (Object k : scheduler.getContext().keySet()) {
				Object v = scheduler.getContext().get(k);
				log.info("  "+k.toString()+" - "+v.toString());
			}
		} catch (SchedulerException ex) { /* ignore */ }
		SchedulerFactory factory = DirectSchedulerFactory.getInstance();
		try {
			log.info("SchedulerFactory num instances: "+factory.getAllSchedulers().size());
		} catch (SchedulerException ex) { /* ignore */ }
		model.addAttribute("schedulerFactory", factory);
		return "welcome";
	}
}
