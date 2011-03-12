package org.elfwerks.meta.beans;

import org.elfwerks.meta.annotations.MetaComposition;

public class ComplexSingularCompositionBean extends BeanSupport {
	
	DateBean dateBean;

	public ComplexSingularCompositionBean() {
		super();
	}
	
	@MetaComposition
	public DateBean getDateBean() { return dateBean; }
	public void setDateBean(DateBean db) { dateBean = db; }
	
}
