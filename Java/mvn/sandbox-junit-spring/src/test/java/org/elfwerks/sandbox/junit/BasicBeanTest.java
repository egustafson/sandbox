package org.elfwerks.sandbox.junit;

import javax.annotation.Resource;

import org.junit.runner.RunWith;
import org.junit.Test;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import static org.junit.Assert.*;

@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration  // default context:  BasicBeanTest-context.xml
public class BasicBeanTest {
	
	@Test
    public void trivialTest() {
        assertTrue( true );
    }
	
	@Resource
	protected Bean testBean;  // configured in the Spring Context
	
	@Test
	public void testTheBean() {
		assertEquals(testBean.getName(), "Bean-Name");
	}
	
}
