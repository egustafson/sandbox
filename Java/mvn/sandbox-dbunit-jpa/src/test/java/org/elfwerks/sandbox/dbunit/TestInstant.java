package org.elfwerks.sandbox.dbunit;

import java.util.Date;

import javax.annotation.Resource;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import javax.sql.DataSource;

import org.dbunit.dataset.DefaultDataSet;
import org.dbunit.dataset.IDataSet;
import org.dbunit.operation.DatabaseOperation;
import org.elfwerks.unittest.AbstractJunit4DatabaseTestCase;
import org.joda.time.DateTimeZone;
import org.joda.time.LocalDate;
import org.joda.time.LocalTime;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.test.annotation.Rollback;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.transaction.annotation.Transactional;

@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration(locations="PersistenceTests-context.xml")
public class TestInstant extends AbstractJunit4DatabaseTestCase {

	@Resource(name="dataSource")
	DataSource dataSource;
	
	@PersistenceContext
	EntityManager em;
	
	@Override
	protected IDataSet getDataSet() throws Exception {
		return new DefaultDataSet();
	}

	@Override
	protected DataSource getDataSource() {
		return dataSource;
	}

	@Override
	protected DatabaseOperation getTearDownOperation() {
		return DatabaseOperation.NONE;
	}
	
	@Test
	@Rollback(false)
	@Transactional
	public void testCreateInstant() throws Exception {
		Date now = new Date();
		Instant i = new Instant();
		i.setTimezone(DateTimeZone.getDefault());
		i.setDate(LocalDate.fromDateFields(now));
		i.setTime(new LocalTime());
		
		em.persist(i);
	}


}
