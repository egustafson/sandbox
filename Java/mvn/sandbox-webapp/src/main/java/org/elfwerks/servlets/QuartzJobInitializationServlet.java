package org.elfwerks.servlets;

import java.text.ParseException;
import java.util.Date;
import javax.servlet.Servlet;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;

import org.quartz.CronTrigger;
import org.quartz.JobDetail;
import org.quartz.ee.servlet.QuartzInitializerServlet;
import org.quartz.SchedulerException;
import org.quartz.SimpleTrigger;
import org.quartz.SchedulerFactory;
import org.quartz.TriggerUtils;

import org.elfwerks.quartz.jobs.HelloJob;

public class QuartzJobInitializationServlet extends HttpServlet implements Servlet {

	private static final long serialVersionUID = 3770660579612340571L;

	@Override
	public void init() throws ServletException {
		try {
			JobDetail job = new JobDetail("job1", "group1", HelloJob.class);

			Date runTime = TriggerUtils.getEvenMinuteDate(new Date());
			SimpleTrigger trigger2 = new SimpleTrigger("trigger1", "group1", runTime);
			CronTrigger trigger = new CronTrigger("trigger1", "group1", "0 0/15/30/45 * * * ?");

			SchedulerFactory factory = (SchedulerFactory)getServletContext().getAttribute(QuartzInitializerServlet.QUARTZ_FACTORY_KEY);
			factory.getScheduler().scheduleJob(job, trigger);
		} catch (SchedulerException ex) {
			throw new ServletException("Failed to schedule job.", ex);
		} catch (ParseException ex) {
			throw new ServletException("Cron expression failed to parse.", ex);
		}
	}

	
	
}
