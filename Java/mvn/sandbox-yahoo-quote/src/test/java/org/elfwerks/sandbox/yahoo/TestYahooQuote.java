package org.elfwerks.sandbox.yahoo;

import static org.junit.Assert.*;

import java.util.List;

import org.joda.time.LocalDate;
import org.junit.Test;

public class TestYahooQuote {

	@Test
	public void testLoad() {
		YahooQuote yq = new YahooQuote();
		String symbol = "INTC";
		LocalDate fromDate = new LocalDate(2007, 7, 6); // July 6, 2007
		LocalDate toDate   = new LocalDate(2008, 2, 5); // February 5, 2008
		List<String> results = yq.loadHistory(symbol, fromDate, toDate);
		for (String line : results) {
			System.out.println(line);
		}
	}
	
	@Test
	public void testLoadRecords() {
		YahooQuote yq = new YahooQuote();
		String symbol = "INTC";
		LocalDate fromDate = new LocalDate(2007, 7, 1); 
		LocalDate toDate   = new LocalDate(2007, 7, 14);
		List<YahooHistoryRecord> records = yq.loadHistoryRecords(symbol, fromDate, toDate);
		for (YahooHistoryRecord r : records) {
			System.out.println(r.toString());
		}
	}
	
	@Test
	public void testQuote() {
		YahooQuote yq = new YahooQuote();
		String symbol = "MSFT";
		List<String> results = yq.quote(symbol, "x");
		for (String line : results) {
			System.out.println(line);
		}
	}
	
	@Test
	public void testDividends() {
		YahooQuote yq = new YahooQuote();
		String symbol = "IBM";
		LocalDate fromDate = new LocalDate(1800, 1, 1);
		LocalDate toDate   = new LocalDate(2009, 1, 1);
		List<String> results = yq.loadDividends(symbol, fromDate, toDate);
		for (String line : results) {
			System.out.println(line);
		}
	}
}
