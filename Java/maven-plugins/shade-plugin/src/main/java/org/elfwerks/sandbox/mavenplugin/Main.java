package org.elfwerks.sandbox.mavenplugin;

import org.joda.time.DateTime;

public class Main {

  public static void main(String[] args) {
    DateTime dt = new DateTime();
    String now = dt.toString();
    System.out.println("It is now: "+now);
  }

}
