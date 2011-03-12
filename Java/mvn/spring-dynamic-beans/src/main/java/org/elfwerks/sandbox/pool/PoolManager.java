package org.elfwerks.sandbox.pool;

import java.util.LinkedList;
import java.util.List;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class PoolManager {
	private final Log log = LogFactory.getLog(this.getClass());

	private final List<PoolElement> poolObjects = new LinkedList<PoolElement>();
	
	public void register(PoolElement e) {
		if ( e != null ) {
			poolObjects.add(e);
			log.debug("Added element["+e.getName()+"]");
		}
	}
	
	public void doProcessing() {
		log.info("Starting pool processing.");
		for (PoolElement e : poolObjects) {
			e.process();
		}
		log.info("Pool processing complete.");
	}
}
