package org.elfwerks.sandbox.yahoo;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.MalformedURLException;
import java.net.URL;
import java.util.LinkedList;
import java.util.List;

import org.joda.time.LocalDate;

public class YahooQuote {

	private final static String urlPrefix = "http://ichart.finance.yahoo.com/table.csv";
	private final static String quotePrefix = "http://finance.yahoo.com/d/quotes.csv?s=";
	
	public List<String> loadHistory(String symbol, LocalDate fromDate, LocalDate toDate) {
		String urlString = formHistoricalUrl(symbol, fromDate, toDate) + "&g=d";
		List<String> response = loadCSV(urlString);
		return response;
	}
	
	public List<YahooHistoryRecord> loadHistoryRecords(String symbol, LocalDate fromDate, LocalDate toDate) {
		List<String> lines = loadHistory(symbol, fromDate, toDate);
		List<YahooHistoryRecord> rows = new LinkedList<YahooHistoryRecord>();
		for (String line : lines) {
			YahooHistoryRecord row = new YahooHistoryRecord(symbol, line);
			rows.add(row);
		}
		return rows;
	}
	
	public List<String> loadDividends(String symbol, LocalDate fromDate, LocalDate toDate) {
		String urlString = formHistoricalUrl(symbol, fromDate, toDate) + "&g=v";
		List<String> response = loadCSV(urlString);
		return response;
	}
	
	public List<String> quote(String symbol, String parameters) {
		String urlString = quotePrefix + symbol + "&f=s" + parameters;
		List<String> response = loadCSV(urlString);
		return response;
	}
	
	
	private List<String> loadCSV(String urlString) {
		List<String> response = new LinkedList<String>();
		try {
			URL url = new URL(urlString);
			InputStream is = url.openStream();
			BufferedReader rd = new BufferedReader(new InputStreamReader(is));
			String line = rd.readLine();
			if (line != null && line.startsWith("Date") ) line = rd.readLine();
			while ( line != null ) {
				response.add(line);
				line = rd.readLine();
			}
		} catch (MalformedURLException ex) {
			// TODO Auto-generated catch block
			ex.printStackTrace();
		} catch (IOException ex) {
			// TODO Auto-generated catch block
			ex.printStackTrace();
		}
		return response;
	}
	
	private String formHistoricalUrl(String symbol, LocalDate fromDate, LocalDate toDate) {
		String urlString = urlPrefix + "?s=" + symbol +
		                   "&a=" + (fromDate.getMonthOfYear()-1) +
		                   "&b=" +  fromDate.getDayOfMonth() +
		                   "&c=" +  fromDate.getYear() +
		                   "&d=" + (toDate.getMonthOfYear()-1) +
		                   "&e=" +  toDate.getDayOfMonth() +
		                   "&f=" +  toDate.getYear();
		return urlString;
	}
}
