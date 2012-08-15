package org.elfwerks.sandbox.scala

import org.elfwerks.sandbox.java.JavaStub;

object Main {
  
    def main(args: Array[String]) = {
    	println("Hello Scala World.")
    	println("from scala-space, the Java message: " + JavaStub.getMessage())
    	JavaStub.sayScalaMessage()
    	println("done.")
    }
	
	def getScalaMesssage():String = "ScalaSpace message"
	
}