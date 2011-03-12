package org.elfwerks.sandbox.spring3.webapp;

import java.io.Serializable;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class PaginationBean implements Serializable {
	/* This class needs to be Serializable IIF it is used as a Session
	 * scoped object.  Either through storage in the HTTPSession object,
	 * or as a Spring session scoped bean.
	 * 
	 * Alternatively, Tomcat could be reconfigured in a non-default way
	 * to not persist sessions across restarts.  (I think?)
	 */
	private static final long serialVersionUID = -4095274213058554937L;

	private final Log log = LogFactory.getLog(this.getClass());

	private String baseName;
	private int pageNumber = 0;

	public PaginationBean() {
		log.debug("Pagination Bean Constructed.");
	}
	
	public String getBaseName() { return baseName; }
	public void setBaseName(String baseName) { this.baseName = baseName; }
	
	public int getPageNumber() { return pageNumber;	}
	public void setPageNumber(int pageNumber) { this.pageNumber = pageNumber; }
	
}
