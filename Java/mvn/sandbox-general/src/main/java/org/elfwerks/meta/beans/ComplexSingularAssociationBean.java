package org.elfwerks.meta.beans;

import org.elfwerks.meta.annotations.MetaAssociation;

public class ComplexSingularAssociationBean extends BeanSupport {

	DateBean dateBean;

	public ComplexSingularAssociationBean() {
		super();
	}
	
	@MetaAssociation
	public DateBean getDateBean() { return dateBean; }
	public void setDateBean(DateBean db) { dateBean = db; }
	

}
