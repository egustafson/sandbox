package org.elfwerks.sandbox.dbunit;

import javax.sql.DataSource;

import org.springframework.jdbc.datasource.SingleConnectionDataSource;
import org.springframework.jdbc.support.incrementer.DataFieldMaxValueIncrementer;
import org.springframework.jdbc.support.incrementer.H2SequenceMaxValueIncrementer;

public class H2TestingUtil {// don't instantiate
	private H2TestingUtil() { }
	/*
	 * Create a new H2 in-process in-memory database
	 */
	@SuppressWarnings("deprecation")
	public static DataSource getNewDataSource() {
		DataSource ds = new SingleConnectionDataSource("org.h2.Driver", "jdbc:h2:mem:", "sa", "", true);
		return ds;
	}

	/**
	 * Creates a DataFieldMaxValueIncrementer tied to a h2 driver and sequence (Spring
	 * @param ds the DataSource the incrementer uses
	 * @return an incrementer to use
	 */
	public static DataFieldMaxValueIncrementer getNewIncrementer(final DataSource ds, final String sequenceName) {
		return new H2SequenceMaxValueIncrementer(ds, sequenceName);
	}

	/**
	 * Creates the tables in the database
	 * @param ds the DataSource to CREATE into
	 */
	public static void createTables(final DataSource ds) {
		try {
			final String stmt = "create table USERS(username varchar(50), email varchar(50), Primary Key (username));";
			ds.getConnection().prepareStatement(stmt).execute();
		} catch (Exception e) {
			// do nothing -- probably just called twice as part of a unit tests
		}
	}
}