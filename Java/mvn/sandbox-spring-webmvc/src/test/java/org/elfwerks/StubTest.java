package org.elfwerks;

import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;

public class StubTest extends TestCase {

    public StubTest( String testName ) {
        super( testName );
    }

    public static Test suite() {
        return new TestSuite( StubTest.class );
    }

    public void testApp() {
        assertTrue( true );
    }
}
