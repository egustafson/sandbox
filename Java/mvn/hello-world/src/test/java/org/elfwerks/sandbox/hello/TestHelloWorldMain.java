package org.elfwerks.sandbox.hello;

import static org.junit.Assert.*;

import org.junit.Before;
import org.junit.Test;

public class TestHelloWorldMain {
	
	private static String validResponse = "Hello, world.";
	
	private HelloWorldMain hwObject;
	
	@Before
	public void setUp() {
		hwObject = new HelloWorldMain();
	}
	
	@Test
	public void testHelloMsg() {
		String msg = hwObject.helloMsg();
		assertEquals(msg, validResponse);
	}

}
