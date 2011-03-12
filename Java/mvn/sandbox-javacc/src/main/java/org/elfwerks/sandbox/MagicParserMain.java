package org.elfwerks.sandbox;

import java.io.ByteArrayInputStream;

import org.elfwerks.parsers.magic.*;

public class MagicParserMain {

    public static void main(String[] args) {
        String input = "17 0x11 021";
        ByteArrayInputStream bais = new ByteArrayInputStream(input.getBytes());
        
        MagicParser p = new MagicParser(bais);
        
        try {
            int value;
            value = p.number();
            value = p.number();
            value = p.number();
            System.out.println(value);
        } catch (ParseException ex) {
            // TODO Auto-generated catch block
            ex.printStackTrace();
        }
    }

}
