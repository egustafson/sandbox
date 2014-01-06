package org.elfwerks.sandbox.jetty;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.joda.time.DateTime;
import org.joda.time.format.DateTimeFormatter;
import org.joda.time.format.ISODateTimeFormat;

/**
 * Servlet implementation class NullServlet
 */
public class NullServlet extends HttpServlet {
	private static final long serialVersionUID = 1L;
	
	private static final DateTimeFormatter isoFmt = ISODateTimeFormat.dateTime();
       
	/**
	 * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse response)
	 */
	protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
        DateTime now = new DateTime();
        response.setStatus(200);
        response.getOutputStream().print("Now: "+ isoFmt.print(now));
	    response.getOutputStream().close();
	}

}
