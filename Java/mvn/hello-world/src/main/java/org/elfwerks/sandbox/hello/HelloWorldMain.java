package org.elfwerks.sandbox.hello;

public class HelloWorldMain {

	public static void main(String[] args) {
		HelloWorldMain hw = new HelloWorldMain();
		String msg = hw.helloMsg();
		System.out.println(msg);
	}
	
	public String helloMsg() {
		return "Hello, world.";
	}
	
}
