package org.elfwerks.meta.beans;

public class BeanMetaException extends Exception {

	private static final long serialVersionUID = 4854441463193589503L;
	
	private final Bean bean;

	/**
	 * 
	 */
	public BeanMetaException(Bean b) {
		super();
		bean = b;
	}

	/**
	 * @param message
	 * @param cause
	 */
	public BeanMetaException(String message, Throwable cause, Bean b) {
		super(message, cause);
		bean = b;
	}

	/**
	 * @param message
	 */
	public BeanMetaException(String message, Bean b) {
		super(message);
		bean = b;
	}

	/**
	 * @param cause
	 */
	public BeanMetaException(Throwable cause, Bean b) {
		super(cause);
		bean = b;
	}

	public Bean getBean() {
		return bean;
	}
	
}
