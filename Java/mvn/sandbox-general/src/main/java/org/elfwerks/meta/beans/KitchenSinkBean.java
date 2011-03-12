package org.elfwerks.meta.beans;

import java.util.Date;
import java.util.Set;
import java.util.HashSet;


import org.elfwerks.meta.annotations.MetaAssociation;
import org.elfwerks.meta.annotations.MetaComposition;
import org.elfwerks.meta.annotations.MetaProperty;

public class KitchenSinkBean extends BeanSupport {

	Set<DateBean> dates = new HashSet<DateBean>();
	Set<IntegerBean> numbers = new HashSet<IntegerBean>();
	StringBean value;
	String name;
	Integer altId;
	Date timestamp = new Date();
	
	public KitchenSinkBean() {
		super();
	}
	
	@MetaProperty
	public Integer getAltId() { return altId; }
	public void setAltId(Integer ai) { altId = ai; }
	
	@MetaProperty
	public String getName() { return name; }
	public void setName(String n) { name = n; }
	
	@MetaProperty
	public Date getTimestamp() { return timestamp; }
	public void setTimestamp(Date ts) { timestamp = ts; }
	
	@MetaComposition
	public StringBean getValue() { return value; }
	public void setValue(StringBean v) { value = v; }
	
	@MetaComposition
	public Set<IntegerBean> getNumbers() { return numbers; }
	public void setNumbers(Set<IntegerBean> n) { numbers = n; }
	public void addNumber(Integer n) {
		IntegerBean ib = new IntegerBean();
		ib.setValue(n);
		numbers.add(ib);
	}
	
	@MetaAssociation
	public Set<DateBean> getDates() { return dates; }
	public void setDates(Set<DateBean> d) { dates = d; }
	public void addDate(DateBean d) {
		dates.add(d);
	}
	public void addDate(Date d) {
		DateBean db = new DateBean();
		db.setValue(d);
		addDate(db);
	}
	
}
