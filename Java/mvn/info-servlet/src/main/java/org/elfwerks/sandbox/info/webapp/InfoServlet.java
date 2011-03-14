package org.elfwerks.sandbox.info.webapp;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Enumeration;
import java.util.Properties;
import java.util.Set;
import java.util.SortedSet;
import java.util.TreeSet;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

/**
 * Servlet implementation class InfoServlet
 */
public class InfoServlet extends HttpServlet {
	private static final long serialVersionUID = 1L;
	private final Log log = LogFactory.getLog(this.getClass());
       
    public InfoServlet() {
        super();
    }

	/**
	 * @see Servlet#init(ServletConfig)
	 */
	public void init(ServletConfig config) throws ServletException {
		log.info("Servlet initialized.");
	}

	/**
	 * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse response)
	 */
	protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
	    PrintWriter wr = response.getWriter();
	    Properties sysProps = System.getProperties();
	    @SuppressWarnings("rawtypes")
        Enumeration keys = sysProps.keys();
	    SortedSet<String> sortedKeys = new TreeSet<String>(); 
	    while (keys.hasMoreElements()) {
	        sortedKeys.add((String)keys.nextElement());
	    }
	    for (String key : sortedKeys) {
            wr.println(key+" = "+sysProps.getProperty(key));
	    }
	}

}
