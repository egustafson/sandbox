package org.elfwerks.sandbox.dbunit;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import static org.junit.Assert.*;

@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration(locations="PersistenceTests-context.xml")
public class TestPersistenceInitialization {

	@Test
	public void TestPersistenceInit() {
		assertTrue(true);
	}
	
}
