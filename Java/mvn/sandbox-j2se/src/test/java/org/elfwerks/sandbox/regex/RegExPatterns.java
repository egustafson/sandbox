package org.elfwerks.sandbox.regex;

import java.util.regex.Pattern;

import org.junit.Test;
import static org.junit.Assert.*;


public class RegExPatterns {

    @Test
    public void testHtmlEscapeMatch() {
        String pat = ".*(\\?|&)Signature=([^\\s]*((%2F)|(%2B))[^\\s]*)+$";
        Pattern htmlPat = Pattern.compile(pat);
        
        String t1 = "http://foo:80/path?P1=foo&Signature=asdfg%2Afds%2Fasd";
        assertTrue(htmlPat.matcher(t1).matches());
        
        String t2 = "x&Signature=xx%2E%2Fxyz%2Bxx";
        assertTrue(htmlPat.matcher(t2).matches());
    }
    
    @Test
    public void testCharacterClass() {
    	String re = "([\\w\\p{Punct}&&[^\\-]]+)";
    	Pattern pat = Pattern.compile(re);
    	
    	String t1 = "abcdefghijklmnopqrstuvwxyz0123456789~!@#$%^&*()_+:;|";
    	assertTrue(pat.matcher(t1).matches());
    	
    	String t2 = "a-b";
    	assertFalse(pat.matcher(t2).matches());
    }
    
    @Test
    public void testDateMatch() {
    	String dates[] = { "junk 2010-08-15T17:00:00 ", " 2010-08-15T19:23:17 " };
    	Pattern p = Pattern.compile("2010-08-15T1[7-9]");
    	for (String d : dates) {
    		assertTrue(p.matcher(d).find());
    	}
    }
    
    @Test
    public void testAnotherDateMatch() {
    	int y = 2010;  int m = 8; int d = 15;
		String d1 = String.format("%4d-%02d-%02d", y, m, d-1);
		String d2 = String.format("%4d-%02d-%02d", y, m, d);
		assertTrue(Pattern.compile(d1+"T1[7-9]").matcher(" 2010-08-14T18:23:17 ").find());
		assertTrue(Pattern.compile(d1+"T2").matcher(" 2010-08-14T23:23:17 ").find());
		assertTrue(Pattern.compile(d2+"T0").matcher(" 2010-08-15T03:23:17 ").find());
		assertTrue(Pattern.compile(d2+"T1[0-6]").matcher(" 2010-08-15T12:23:17 ").find());

    	assertFalse(Pattern.compile(d1+"T1[7-9]").matcher(" 2010-08-14T16:23:17 ").find());
		assertFalse(Pattern.compile(d1+"T2").matcher(" 2010-08-15T23:23:17 ").find());
		assertFalse(Pattern.compile(d2+"T0").matcher(" 2010-08-14T03:23:17 ").find());
		assertFalse(Pattern.compile(d2+"T1[0-6]").matcher(" 2010-08-15T17:23:17 ").find());
    }
    
}
