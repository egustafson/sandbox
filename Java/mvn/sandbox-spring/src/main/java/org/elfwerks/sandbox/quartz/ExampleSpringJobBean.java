package org.elfwerks.sandbox.quartz;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.quartz.JobExecutionContext;
import org.quartz.JobExecutionException;
import org.springframework.scheduling.quartz.QuartzJobBean;

public class ExampleSpringJobBean extends QuartzJobBean {

	private static final Log log = LogFactory.getLog(ExampleSpringJobBean.class);
	
	private long sleepMillis;
	
	public void setSleepMillis(long millis) {
		sleepMillis = millis;
	}
	
	@Override
	protected void executeInternal(JobExecutionContext context) throws JobExecutionException {
		log.info("Sleeping for "+sleepMillis+"ms.");
		try {
			Thread.sleep(sleepMillis);
		} catch (InterruptedException ex) {
			log.info("Awoke early from sleep, caught InterruptedException.");
		}
		log.info("Done sleeping, completing job.");
	}

}
