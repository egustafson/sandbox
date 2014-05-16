package org.elfwerks.groovy
import org.elfwerks.groovy.GroovyMain;
import org.junit.Test
import org.junit.Assert

class GroovyTest {

	@Test
	void testMethod() {
		GroovyMain.main null
		Assert.assertTrue true
	}
}