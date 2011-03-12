package org.elfwerks;

import javax.servlet.ServletException;

/**
 * Servlet implementation class for Servlet: JSPSandboxServlet
 *
 */
 public class JSPSandboxServlet extends javax.servlet.http.HttpServlet implements javax.servlet.Servlet {
   static final long serialVersionUID = 1L;
	
	/**
	 * @see javax.servlet.GenericServlet#init()
	 */
	public void init() throws ServletException {
		getServletConfig().getServletContext().setAttribute("demo", "'demo value'");
	}   
}