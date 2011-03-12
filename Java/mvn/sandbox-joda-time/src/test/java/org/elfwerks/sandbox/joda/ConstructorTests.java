package org.elfwerks.sandbox.joda;

import static org.junit.Assert.*;

import org.joda.time.DateTime;
import org.joda.time.DateTimeZone;
import org.joda.time.format.DateTimeFormat;
import org.joda.time.format.DateTimeFormatter;
import org.joda.time.format.ISODateTimeFormat;
import org.junit.Test;

public class ConstructorTests {

	private DateTime refDateTime = new DateTime("2010-05-06T10:20:01.001Z", DateTimeZone.UTC);
	
	@Test
	public void testISO8601() {
		String ts = "2010-05-06T10:20:01.001";
		DateTime dt = new DateTime(ts, DateTimeZone.UTC);
		System.out.println(dt);
		assertTrue(refDateTime.equals(dt));
	}
	
	@Test
	public void testBasicISO8601() {
		String ts = "20100506T102001.001Z";
		DateTimeFormatter fmt = ISODateTimeFormat.basicDateTime().withZone(DateTimeZone.UTC); 
		DateTime dt = fmt.parseDateTime(ts);
		System.out.println(dt);
		assertTrue(refDateTime.equals(dt));
	}
	
	@Test
	public void testISO8601_overrideTz() {
		String ts = "2010-05-06T04:20:01.001-0600";
		DateTime dt = new DateTime(ts, DateTimeZone.UTC);
		System.out.println(dt);
		assertTrue(refDateTime.equals(dt));
	}

	public void testBasicISO8601_noTz() {
		String ts = "20100506T102001.001";
		DateTimeFormatter fmt = DateTimeFormat.forPattern("YYYYMMddTHHmmss.SSS").withZone(DateTimeZone.UTC);
		DateTime dt = fmt.parseDateTime(ts);
		System.out.println(dt);
		assertTrue(refDateTime.equals(dt));
	}
	
	@Test
	public void testCompact() {
		String ts = "20100506102001.001";
		DateTimeFormatter fmt = DateTimeFormat.forPattern("YYYYMMddHHmmss.SSS").withZone(DateTimeZone.UTC);
		DateTime dt = fmt.parseDateTime(ts);
		System.out.println(dt);
		assertTrue(refDateTime.equals(dt));
	}
	
	@Test(expected=IllegalArgumentException.class)
	public void testBadISO8601() {
		String ts = "20100506102001.001";
		DateTime dt = new DateTime(ts, DateTimeZone.UTC);
		System.out.println(dt);
	}
	
	@Test(expected=IllegalArgumentException.class)
	public void testBadCompactFormat() {
		String ts = "20100506T102001.001";
		DateTimeFormatter fmt = DateTimeFormat.forPattern("YYYYMMddHHmmss.SSS").withZone(DateTimeZone.UTC);
		DateTime dt = fmt.parseDateTime(ts);
		System.out.println(dt);
	}
	
	
}
