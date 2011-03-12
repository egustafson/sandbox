package org.elfwerks.sandbox.rtcomp;

import javax.annotation.PostConstruct;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Required;

public class SimpleComponent implements RuntimeComponent, SimpleComponentMBean {
	private final Log log = LogFactory.getLog(this.getClass());

	private String name;
	private Registrar registrar;
	
	@Required
	public void setName(String n) { name = n; }
	@Override
	public String getName() { return name; }
	
	@Autowired
	public void setRegistrar(Registrar r) { registrar = r; }
	
	@PostConstruct
	public void afterPropertiesSet() {
		speak("afterPropertiesSet()");
		registrar.register(this);
	}
	
	@Override
	public void speak(String message) {
		log.info("["+name+"] - "+message);
	}
	
	@Override
	public String toString() { return name; }

}
