package org.elfwerks.meta.beans;

import org.elfwerks.meta.annotations.MetaId;

public class BeanSupport extends Bean {

	int id;
	
	BeanSupport() {
		setId(Bean.getNextId());
	}
	
	@MetaId
	public int getId() { return id; }
	public void setId(int i) {
		if ( id > 0 ) {
			Bean.removeBean(id);
		}
		id = i;
		try { 
			Bean.registerBean(this);
		} catch (BeanMetaException ex) {
			throw new RuntimeException("(INTERNAL ERROR)", ex);
		}
	}
	
	
}
