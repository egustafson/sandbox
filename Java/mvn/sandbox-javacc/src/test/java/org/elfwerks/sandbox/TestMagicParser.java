package org.elfwerks.sandbox;

import static org.junit.Assert.*;

import java.io.ByteArrayInputStream;

import org.junit.Test;

import org.elfwerks.parsers.magic.MagicParser;
import org.elfwerks.parsers.magic.ParseException;

public class TestMagicParser {

    MagicParser p;
    
    private void init(String input) {
        ByteArrayInputStream bais = new ByteArrayInputStream(input.getBytes());
        p = new MagicParser(bais);
    }
    
    @Test
    public void testNumber() throws ParseException {
        init("17 0x11 021");
        assertEquals(p.number(), 17);
        assertEquals(p.number(), 17);
        assertEquals(p.number(), 17);
    }
    
    @Test
    public void testLevel() throws ParseException {
        init("> >> >>> >>>> >>>>> >>>>>>");
        /* level 0 has to be tested at a higher level in the parser */
        assertEquals(p.level(), 1);
        assertEquals(p.level(), 2);
        assertEquals(p.level(), 3);
        assertEquals(p.level(), 4);
        assertEquals(p.level(), 5);
        assertEquals(p.level(), 6);
    }

    @Test
    public void testOffset() throws ParseException {
        init("1234 0x1ff 0733");
        assertEquals(p.offset(), 1234);
        assertEquals(p.offset(), 0x1ff);
        assertEquals(p.offset(), 0733);  /* octal */
    }
    
    @Test
    public void testSearch() throws ParseException {
        init("string belong beshort byte byte&0x7f beshort&0177 belong&9876");
        p.search();
        p.search();
        p.search();
        p.search();
        p.search();
        p.search();
        p.search();
    }
    
}
