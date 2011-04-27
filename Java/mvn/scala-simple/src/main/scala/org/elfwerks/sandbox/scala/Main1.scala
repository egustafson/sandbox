package org.elfwerks.sandbox.scala

object Main1 {

  def main(args: Array[String]): Unit = {  
      
      val p  = new Point(1,1)
      val p2 = new Point(1,2)
      val p3 = new Point3d(3,2,1)
      println(p)
      println(p2.toString)
      println(p3)
  }

}