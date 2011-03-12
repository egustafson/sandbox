package org.elfwerks.sandbox.quartz;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class ExampleJobPerformingBean {
	private final Log log = LogFactory.getLog(this.getClass());

	public void doJob() {
		log.info("Performed job.");
	}
	
}
