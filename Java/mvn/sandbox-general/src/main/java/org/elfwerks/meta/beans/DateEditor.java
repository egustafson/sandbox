package org.elfwerks.meta.beans;

import java.beans.PropertyEditorSupport;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

/**
 * A custom java.util.Date PropertyEditor which implements the textual
 * representation as the XML Schema xs:dateTime formatting. (aka: ISO 8601)
 */
public class DateEditor extends PropertyEditorSupport {

	/** XML Schema xs:dateTime format */
	private static String formatString = "yyyy-MM-dd'T'HH:mm:ssZ"; 
	private static SimpleDateFormat dateFormat = new SimpleDateFormat(formatString); 
	
	private Date date;

	@Override
	public Object getValue() {
		return date.clone();
	}

	@Override
	public void setValue(Object value) {
		date = (Date)((Date)value).clone();
	}
	
	@Override
	public void setAsText(String text) throws IllegalArgumentException {
		try {
			date = dateFormat.parse(text);
		} catch (ParseException ex) {
			throw new IllegalArgumentException(ex);
		}
	}

	@Override
	public String getAsText() {
		return dateFormat.format(date);
	}
	
}
