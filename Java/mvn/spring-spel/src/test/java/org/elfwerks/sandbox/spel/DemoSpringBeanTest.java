package org.elfwerks.sandbox.spel;

import static org.junit.Assert.*;

import javax.annotation.Resource;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration // default context:  DemoSpringBeanTest-context.xml
public class DemoSpringBeanTest {

    @Resource
    protected DemoSpringBean testBean;
    
    @Test
    public void testGetName() {
        assertEquals("Configuration Name", testBean.getName());
    }

    @Test
    public void testGetVersion() {
        assertEquals("1.2.0", testBean.getVersion());
    }

    @Test
    public void testGetPort() {
        assertEquals(2120, testBean.getPort());
    }

    @Test
    public void testToString() {
        String str = testBean.toString();
        System.out.println(str);
        assertNotNull(str);
    }

}
