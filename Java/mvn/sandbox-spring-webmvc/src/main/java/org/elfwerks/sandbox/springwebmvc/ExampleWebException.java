package org.elfwerks.sandbox.springwebmvc;

public class ExampleWebException extends Exception {
	private static final long serialVersionUID = -7630478779392439755L;

	public ExampleWebException(String message, Throwable cause) {
		super(message, cause);
	}

	public ExampleWebException(String message) {
		super(message);
	}

	
}
