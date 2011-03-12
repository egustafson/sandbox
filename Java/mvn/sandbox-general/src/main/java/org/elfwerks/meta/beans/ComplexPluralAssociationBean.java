package org.elfwerks.meta.beans;

import java.util.Set;
import java.util.TreeSet;

import org.elfwerks.meta.annotations.MetaAssociation;


public class ComplexPluralAssociationBean extends BeanSupport {

	Set<DateBean> dates = new TreeSet<DateBean>();

	public ComplexPluralAssociationBean() {
		super();
	}
	
	@MetaAssociation
	public Set<DateBean> getDates() { return dates; }
	public void setDates(Set<DateBean> d) { dates = d; }
	
	public void setDate(DateBean d) { dates.add(d); }

}
