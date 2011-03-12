package org.elfwerks.sandbox.dbunit;

import java.util.List;

import javax.sql.DataSource;

import org.dbunit.DataSourceBasedDBTestCase;
import org.dbunit.dataset.IDataSet;
import org.dbunit.dataset.xml.FlatXmlDataSet;
import org.dbunit.operation.DatabaseOperation;

public class JdbcUserDaoTest extends DataSourceBasedDBTestCase {
	private JdbcUserDao instance;
	private DataSource dataSource;
	public JdbcUserDaoTest(String testName) {
		super(testName);
	}
	

	public void testGetAllUsers() {
		List<User> users = instance.getAllUsers();
		assertEquals(2, users.size());
		// more testing ...
	}

	
	/**
	 * setUp is called by JUnit before every test method is run
	 */
	protected void setUp() throws Exception {
		// create the database
		dataSource = H2TestingUtil.getNewDataSource();
		H2TestingUtil.createTables(dataSource);
		// defer to dbunit
		super.setUp();
		// create the instance to test
		instance = new JdbcUserDao();
		instance.setDataSource(dataSource);
	}

	/**
	 * tearDown is called by JUnit after each test and needs to be present for the data to be cleared
	 */
	protected void tearDown() throws Exception {
		super.tearDown();
	}
	
	/**
	 * DBUnit will insert into the dataSource by retrieving it through method
	 */
	protected DataSource getDataSource() {
		return dataSource;
	}
	
	/**
	 * DBUnit calls this to get the data to inject to the database
	 */
	protected IDataSet getDataSet() throws Exception {
		return new FlatXmlDataSet(getClass().getResource("UserTestData.xml"));
	}
	
	/* What DbUnit does with the existing data and with data in the XML file */ 
	protected DatabaseOperation getSetUpOperation() { 
		return DatabaseOperation.CLEAN_INSERT; 
	} 
	
	/* What DbUnit does with the data after the test is done. */ 
	protected DatabaseOperation getTearDownOperation() throws Exception { 
		return DatabaseOperation.DELETE_ALL; 
	}
}