package org.elfwerks.sandbox.rtcomp;

import java.util.LinkedList;
import java.util.List;

import javax.management.ObjectName;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jmx.export.MBeanExporter;

public class DaemonHolder implements Registrar, Runnable, DaemonHolderMBean {
	private final Log log = LogFactory.getLog(this.getClass()); 

	private Boolean shutdownFlag = false;
	
	@Autowired
	private MBeanExporter mbeanExporter;
	
	private List<RuntimeComponent> registeredComponents = new LinkedList<RuntimeComponent>();

	public void register(RuntimeComponent rtc) {
		registeredComponents.add(rtc);
		ObjectName jmxName;
		try {
			jmxName = new ObjectName("org.elfwerks:comp=rtcomp,name="+rtc.toString());
			mbeanExporter.registerManagedResource(rtc, jmxName);
		} catch (Exception ex) {
			log.error("Failed to register ["+rtc.toString()+"] - "+ex.getClass().getSimpleName(), ex);
		}
		log.info("Added runtime-component: "+rtc.toString());
	}
	
	public synchronized void run() {
		log.info("Daemon started.");
			while ( !shutdownFlag ) {
				log.debug("blocking on shutdownFlag");
				try { this.wait();
				} catch (InterruptedException ex) { /* ignore */ }
				log.debug("Awake in run().");
			}
		log.info("Daemon finished.");
	}
	
	public synchronized void shutdown() { 
		log.info("Received shutdown message.");
		shutdownFlag = true;
		this.notifyAll();
	}
	
}
