package org.elfwerks.sandbox.java;

import org.elfwerks.sandbox.scala.Main;

public class JavaStub {

	public static String getMessage() {
		return "message from JavaSpace";
	}
	
	public static void sayScalaMessage() {
		String msg = Main.getScalaMesssage();
		System.out.println("From java-space, the Scala message: " + msg);
	}
}
