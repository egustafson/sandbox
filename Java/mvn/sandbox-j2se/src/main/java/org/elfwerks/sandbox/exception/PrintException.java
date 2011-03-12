package org.elfwerks.sandbox.exception;

import java.io.PrintWriter;
import java.io.StringWriter;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class PrintException {
	private static final Log log = LogFactory.getLog(PrintException.class);

	public static void main(String[] args) {
		log.info("Log Start.");
		try {
			causeException();
		} catch (Exception ex) {
			System.out.println("----- Printing Exception -----"); 
			printException(ex); 
			System.out.println("----- Logging (Log4j) Exception");
			logException(ex);
			System.out.println("----- Printing String buffered Exception");
			String exStackTrace = extractExceptionMsg(ex);
			System.out.print(exStackTrace);
			System.out.println("----- done -----");
		}
		log.info("done.");
	}

	private static void printException(Throwable ex) {
		ex.printStackTrace(System.out);  // defaults to stderr
	}
	
	private static void logException(Throwable ex) {
		log.error("Caught ["+ex.getClass().getSimpleName()+"] cause: "+ex.getMessage(), ex);
	}
	
	private static String extractExceptionMsg(Throwable ex) {
		StringWriter sw = new StringWriter();
		PrintWriter pw = new PrintWriter(sw);
		ex.printStackTrace(pw);
		return sw.toString();
	}
	
	private static void causeException() throws Exception {
		throw new Exception("Demo Exception");
	}
	
}
