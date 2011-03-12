package org.elfwerks.servlets;

import java.io.IOException;

import javax.servlet.Servlet;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class HelloServlet extends HttpServlet implements Servlet {

	private static final long serialVersionUID = 4975706544045928221L;
	private static final Log log = LogFactory.getLog(HelloServlet.class);
	
	@Override
	public void init() throws ServletException {
		String msg = getClass().getSimpleName()+" loaded and initialized.";
		this.getServletContext().log(msg);
		log.info(msg);
	}

	@Override
	protected void doGet(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
		resp.getWriter().println("Hello, world.");
		log.info("Sent hello message.");
	}

}
