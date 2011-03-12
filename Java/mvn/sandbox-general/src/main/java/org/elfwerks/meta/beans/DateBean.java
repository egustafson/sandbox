package org.elfwerks.meta.beans;

import java.util.Date;

import org.elfwerks.meta.annotations.MetaProperty;

public class DateBean extends BeanSupport {

	Date value;

	public DateBean() {
		super();
	}
	
	@MetaProperty(required=true)
	public Date getValue() { return value; }
	public void setValue(Date d) { value = d; }

}
