package org.elfwerks.sandbox.scala.log4j

import org.apache.commons.logging.LogFactory;

object Main {
  val log = LogFactory.getLog(Main.getClass)
  
  def main(args: Array[String]): Unit = {  
      log.info("Log message.")
  }

}