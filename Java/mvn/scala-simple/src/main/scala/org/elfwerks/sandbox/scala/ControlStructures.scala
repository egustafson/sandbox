package org.elfwerks.sandbox.scala

object ControlStructures {

  var assertionEnabled = true
    
  def main(args: Array[String]) {  

      print("Count to 10: ")
      for (ii <- 1 to 10) print(ii + " ")
      println()
      
      listFiles()
      
      listFilesRecursive()

      protoAssert(true == true)
      
  }
  
  def listFiles() {
      val filesHere = (new java.io.File(".")).listFiles
      for ( file <- filesHere 
            if !file.getName.startsWith(".")
            if file.isFile
           ) println(file)
  }
  
  def listFilesRecursive() {
      
  }

  def protoAssert(predicate: => Boolean) {
      if ( assertionEnabled ) {
          println("Assertion tested.")
          if ( !predicate ) throw new AssertionError
      }
  }
  
}