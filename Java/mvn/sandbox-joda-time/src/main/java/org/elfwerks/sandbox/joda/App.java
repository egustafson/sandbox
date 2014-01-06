package org.elfwerks.sandbox.joda;

import org.joda.time.DateTime;
import org.joda.time.format.DateTimeFormatter;
import org.joda.time.format.ISODateTimeFormat;

/* See unit tests for additional use-cases
 */

public class App {
    
    public static void main( String[] args ) {
        DateTimeFormatter isoFmt = ISODateTimeFormat.dateTime();
        DateTime now = new DateTime();
        System.out.println("It is now "+isoFmt.print(now));
    }
    
}
