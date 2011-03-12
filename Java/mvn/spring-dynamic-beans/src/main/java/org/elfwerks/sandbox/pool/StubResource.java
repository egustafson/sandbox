package org.elfwerks.sandbox.pool;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Required;

public class StubResource {
	private final Log log = LogFactory.getLog(this.getClass());
	
	private String name;
	
	@Required
	public void setName(String n) { name = n; }

	public void process(String requestorName) {
		log.info("Resource["+name+"] processing request for: "+requestorName);
	}
}
