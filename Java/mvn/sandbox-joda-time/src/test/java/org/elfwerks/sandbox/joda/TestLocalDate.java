package org.elfwerks.sandbox.joda;

import static org.junit.Assert.*;

import org.joda.time.LocalDate;
import org.junit.Test;


public class TestLocalDate {

	@Test
	public void testToString() {
		String dateFmt = "%04d-%02d-%02d";
		LocalDate today = new LocalDate();
		String isoDate = String.format(dateFmt, today.getYear(), today.getMonthOfYear(), today.getDayOfMonth());
		assertEquals(isoDate, today.toString());
	}
	
	@Test
	public void testSetDate() {
		LocalDate newYear = new LocalDate(2000, 1, 1);
		assertEquals(newYear.getYear(), 2000);
		assertEquals(newYear.getMonthOfYear(), 1);
		assertEquals(newYear.getDayOfMonth(), 1);
	}
	
	@Test(expected=org.joda.time.IllegalFieldValueException.class)
	public void testBogusSetDate() {
		@SuppressWarnings("unused")
		LocalDate bogusDate = new LocalDate(2000, 2, 30);
	}
	
	@Test(expected=org.joda.time.IllegalFieldValueException.class)
	public void testBadLeapYear() {
		@SuppressWarnings("unused")
		LocalDate badLeapYear = new LocalDate(2001, 2, 29);
	}
	
	@Test
	public void testDateArithmetic() {
		LocalDate date = new LocalDate(2000, 1, 1);
		assertEquals(date.minusDays(1).getYear(), 1999);
		assertEquals(date.minusDays(1).getMonthOfYear(), 12);
		assertEquals(date.minusDays(1).getDayOfMonth(), 31);
	}
	
}
