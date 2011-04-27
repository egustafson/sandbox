package org.elfwerks.sandbox.scala

class Point3d(x : Int, y : Int, z : Int) extends Point(x,y) {

    override def toString = "("+x+", "+y+", "+z+")"
}