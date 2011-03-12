package org.elfwerks.log4j;

import java.net.URL;
import java.io.File;

import org.apache.log4j.*;

public class InitLog4j {

    private static final Logger log = Logger.getLogger(InitLog4j.class);
    private static final String defaultConfigfile    = "initlog4j-log4j.properties";
    private static final String defaultLogfile       = "initlog4j-logfile.log";
    private static final String defaultPatternLayout = "%-5p %d {%c}:  %m%n";
    
    /**
     */
    public static void main(String[] args) {
        initLog4j(null, null);
        log.info("Logging started.");
    }

    /**
     * This routine has two goals:
     *   1. Initialize and start log4j using an optionally supplied properties file.
     *   2. Create and add a FileAppender using an optionally provided log file name. 
     * @param configfile the (optional) name of the log4j.properties file to use.
     * @param logfile the (optional) name of the file to create a FileAppender for.
     * @return this method will return on success, and exit the JVM on failure.
     */
    private static void initLog4j(String configfile, String logfile) {
        if ( logfile == null ) {
            logfile = defaultLogfile;
        }
        URL log4jConfig = file2url(configfile);
        if ( log4jConfig == null ) {
            log4jConfig = file2url(defaultConfigfile);
        }
        if ( log4jConfig == null ) {
            log4jConfig = InitLog4j.class.getClassLoader().getResource(defaultConfigfile);
        }
        try {
            PropertyConfigurator.configure(log4jConfig);
            Layout layout = new PatternLayout(defaultPatternLayout);
            FileAppender appender = new FileAppender(layout, logfile);
            Logger.getRootLogger().addAppender(appender);
        } catch (Exception ex) {
            System.err.println("ERROR - Failed to initialize logging.");
            System.err.println("Program terminating, no action performed.");
            System.exit(1);
        }
    }
    
    private static URL file2url(String filename) {
        try {
            File f = new File(filename);
            if ( f.canRead() ) {
                return f.toURI().toURL();
            }
        } catch (Exception ex) { /* do nothing */ }
        return null;
    }
}
