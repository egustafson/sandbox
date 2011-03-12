package org.elfwerks.sandbox;

import java.io.ByteArrayInputStream;

import org.elfwerks.parsers.simple.ParseException;
import org.elfwerks.parsers.simple.SimpleParser;

/**
 * This class demonstrates a very simple use of the SimpleParser.jj gramar.
 */
public class SimpleParserMain {

    public static void main(String[] args) {
        String input = " 1234 \r\n 4321\r\n666 \n";
        ByteArrayInputStream bais = new ByteArrayInputStream(input.getBytes());
        
        SimpleParser p = new SimpleParser(bais);
        
        try {
            while (true) {
                int value = p.line();
                System.out.println("Result:  "+value);
            }
        } catch (ParseException ex) {
            // TODO Auto-generated catch block
            ex.printStackTrace();
        }
    }
}
