package org.elfwerks.sandbox.web.servlet;

import java.io.IOException;
import java.io.PrintWriter;
import javax.servlet.ServletContext;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.context.ApplicationContext;
import org.springframework.web.context.support.WebApplicationContextUtils;

/**
 * This servlet demonstrates how to retrieve the Spring <code>ApplicationContext</code>
 */
 public class PlainServlet extends javax.servlet.http.HttpServlet implements javax.servlet.Servlet {
	private static final long serialVersionUID = 1591494863460990167L;
	
	private ApplicationContext applicationContext;

	/* (non-Javadoc)
	 * @see javax.servlet.GenericServlet#init()
	 */
	public void init() throws ServletException {
		ServletContext sc = getServletContext();
		applicationContext = WebApplicationContextUtils.getRequiredWebApplicationContext(sc);
	}   

	/* (non-Java-doc)
	 * @see javax.servlet.http.HttpServlet#doGet(HttpServletRequest request, HttpServletResponse response)
	 */
	protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
		response.setContentType("text/plain");
		PrintWriter out = response.getWriter();
		out.println("Spring ApplicationContext Bean Names");
		out.println("------------------------------------");
		String[] beanNames = applicationContext.getBeanDefinitionNames();
		for (String name: beanNames) {
			out.println(name);
		}
	}  	  	  	  

 }