package org.elfwerks.meta.beans;

import org.elfwerks.meta.annotations.MetaProperty;

public class StringBean extends BeanSupport {

	String value;

	public StringBean() {
		super();
	}
	
	@MetaProperty(required=true)
	public String getValue() { return value; }
	public void setValue(String v) { value = v; }

}
