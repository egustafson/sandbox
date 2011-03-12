package org.elfwerks.sandbox.dbunit;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

import org.joda.time.DateTimeZone;
import org.joda.time.LocalDate;
import org.joda.time.LocalTime;

@Entity
@Table(name="instants")
public class Instant {

	@Id
	@Column(name="id")
	@GeneratedValue(strategy=GenerationType.AUTO)
	protected long id;

	@Column(name="the_date")
	protected java.sql.Date date;

	@Column(name="the_time")
	protected java.sql.Time time;

	@Column(name="timezone")
	protected String timezone;

	
	public long getId() { return id; }
	protected void setId(long id) { this.id = id; }
	
	public LocalDate getDate() { return (date == null ? null : LocalDate.fromDateFields(date)); }
	public void setDate(LocalDate date) { this.date = (date == null ? null : new java.sql.Date(date.toDateTimeAtStartOfDay().getMillis())); }
	
	public LocalTime getTime() { return (time == null ? null : new LocalTime(time)); }
	public void setTime(LocalTime time) { this.time = (time == null ? null : new java.sql.Time(time.getMillisOfDay())); }
	
	public DateTimeZone getTimezone() { return (timezone == null ? null : DateTimeZone.forID(timezone)); }
	public void setTimezone(DateTimeZone timezone) { this.timezone = (timezone == null ? null : timezone.getID()); }
	
}
