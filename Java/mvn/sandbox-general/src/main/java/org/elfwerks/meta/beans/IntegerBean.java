package org.elfwerks.meta.beans;

import org.elfwerks.meta.annotations.MetaProperty;

public class IntegerBean extends BeanSupport {

	Integer value;

	public IntegerBean() {
		super();
	}
	
	@MetaProperty(required=true)
	public Integer getValue() { return value; }
	public void setValue(Integer v) { value = v; }
	
}
