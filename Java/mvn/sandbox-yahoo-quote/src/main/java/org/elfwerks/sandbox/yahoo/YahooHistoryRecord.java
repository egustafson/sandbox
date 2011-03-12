package org.elfwerks.sandbox.yahoo;

import java.math.BigDecimal;
import java.math.BigInteger;

import org.joda.time.LocalDate;

public class YahooHistoryRecord {

	private String symbol;
	private LocalDate date;
	private BigDecimal open;
	private BigDecimal close;
	private BigDecimal high;
	private BigDecimal low;
	private BigInteger volume;

	public YahooHistoryRecord(String symbol, String yahooLine) {
		String[] fields = yahooLine.split(",");
		String[] dateFields = fields[0].split("-");
		this.symbol = symbol;
		this.date  = new LocalDate(Integer.parseInt(dateFields[0]), 
								   Integer.parseInt(dateFields[1]), 
								   Integer.parseInt(dateFields[2]));
		this.open  = new BigDecimal(fields[1]);
		this.high  = new BigDecimal(fields[2]);
		this.low   = new BigDecimal(fields[3]);
		this.close = new BigDecimal(fields[4]);
		this.volume= new BigInteger(fields[5]);
	}
	
	public String getSymbol() { return symbol; }
	public LocalDate getDate() { return date; }
	public BigDecimal getOpen() { return open; }
	public BigDecimal getClose() { return close; }
	public BigDecimal getHigh() { return high; }
	public BigDecimal getLow() { return low; }
	public BigInteger getVolume() { return volume; }
	
	public String toString() {
		return symbol+"["+date.toString()+"]"+open+"^"+high+"-"+low+">"+close+"v"+volume;
	}
}
