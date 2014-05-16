package org.elfwerks.groovy;

import org.elfwerks.groovy.JavaMain;
import org.junit.Test;
import org.junit.Assert;

public class JavaTest {

	@Test
	public void testMethod() {
		JavaMain.main(new String[] {});
		Assert.assertTrue(true);
	}
}