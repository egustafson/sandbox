package org.elfwerks.argv;

import java.util.Iterator;

public class DemoArgv {
    
    /**
     * This <code>main()</code> program dumps argument and System properties that 
     * start with 'org.elfwerks'.  The intent is to allow the user to experiment
     * with command line parameters to see what the JVM passes in where.
     * 
     * Conclusion:  
     *   'vm options' come before the '-jar' or classname on the command line and
     *     include "-Dproperty.name=value".  'vm options' do NOT appear in the
     *     <code>argv</code> list.  '-Dproperty' options are incorporated into the
     *     System properties list.
     *     
     *   'arguments' must be placed after the '-jar' or classname on the command line
     *     and appear in the <code>argv</code> array.
     *     
     *   '"-D" arguments' that appear after the '-jar' or classname are considered
     *     normal program arguments and are not incorporated into the System 
     *     properties list. 
     * 
     * @param argv Arguments from the command line
     */
    public static void main(String[] argv) {
        System.out.println("Printing 'argv' array:");
        for (int ii = 0; ii < argv.length; ii++) {
            System.out.println("  "+argv[ii]);
        }
        System.out.println("--");
        Iterator<Object> sysPropKeys = System.getProperties().keySet().iterator();
        while (sysPropKeys.hasNext()) {
            String k = (String)sysPropKeys.next();
            if ( k.startsWith("org.elfwerks") ) {
                System.out.println(k+"="+System.getProperty(k));
            }
        }
        System.out.println("Done.");
    }

}
