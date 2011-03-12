package org.elfwerks.quartz.jobs;

import java.util.Date;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import org.quartz.Job;
import org.quartz.JobExecutionContext;
import org.quartz.JobExecutionException;

public class HelloJob implements Job {

	private static final Log log = LogFactory.getLog(HelloJob.class);
	
	@Override
	public void execute(JobExecutionContext arg0) throws JobExecutionException {
		Date now = new Date();
		log.info("Hello from the HelloJob at "+now.toString());
	}

}
