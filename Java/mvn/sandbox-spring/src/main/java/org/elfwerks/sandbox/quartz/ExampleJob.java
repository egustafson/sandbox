package org.elfwerks.sandbox.quartz;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import org.quartz.Job;
import org.quartz.JobExecutionContext;
import org.quartz.JobExecutionException;


public class ExampleJob implements Job {

	private static final Log log = LogFactory.getLog(ExampleJob.class);
	
	@Override
	public void execute(JobExecutionContext jobCtx) throws JobExecutionException {
		log.info("executing quartz job.");
	}

}
