package org.elfwerks.sandbox.pool;

import java.util.Properties;

import org.springframework.beans.factory.annotation.Autowired;

public class PoolElement {
	
	private StubResource resource;
	private String name;
	private int value;
	
	@Autowired
	public void setResource(StubResource resource) { this.resource = resource; }
	public String getName() { return name; }

	public void configure(String name, Properties p) throws PoolElConfigurationException {
		try {
			this.name = name;
			this.value = Integer.parseInt(p.getProperty("prop"));
		} catch (Exception ex) {
			throw new PoolElConfigurationException();
		}
	}

	
	public void process() {
		resource.process(name);
	}
	
}
