package org.elfwerks.filters;

import java.io.IOException;
import java.io.PrintStream;
import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Date;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

/**
 * 
 */
public class AccessLogFilter implements Filter {

	private static final Log log = LogFactory.getLog(AccessLogFilter.class);
	private PrintStream accessLog = System.out;
	private FilterConfig filterConfig;
	
	@Override
	public void init(FilterConfig filterConfig) throws ServletException {
		this.filterConfig = filterConfig;
		log.info("Initialized.");
	}

	@Override
	public void destroy() {
		log.info("Finalized.");
	}

	@Override
	public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain) throws IOException, ServletException {
		log.info("Filter Started.");

		DateFormat df = new SimpleDateFormat("yyyy-MM-dd'T'HH:mm:ss.SSSZ");
		Date now = new Date();
		
		String clientHost = request.getRemoteHost();
		String clientAddr = request.getRemoteAddr();
		String path = "";
		if ( request instanceof HttpServletRequest ) {
			HttpServletRequest httpReq = (HttpServletRequest)request;
			path = httpReq.getContextPath()+httpReq.getServletPath();
			String pathInfo = httpReq.getPathInfo();
			if ( pathInfo != null ) {
				path = path + pathInfo;
			}
			String queryString = httpReq.getQueryString();
			if ( queryString != null ) {
				path = path + queryString;
			}
		}
		long startTime = System.currentTimeMillis();
		
		chain.doFilter(request, response);

		long endTime = System.currentTimeMillis();
		accessLog.println(df.format(now)+": ("+(endTime-startTime)+"ms) - "+clientHost+"("+clientAddr+") -> "+path);
		
		log.info("Filter Finishing.");
	}

	
	
}
