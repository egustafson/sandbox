package org.elfwerks.sandbox.timer;

import java.util.Timer;
import java.util.TimerTask;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class Main {
    private static final Log log = LogFactory.getLog(Main.class);
    
    public static void main(String[] args) {
        log.info("Start.");
        startTimer();
        log.info("Timer started, sleeping for 20s.");
        try {
            Thread.sleep(20*1000);
        } catch (InterruptedException ex) {
            log.error("Caught InterruptedException", ex);
        }
        log.info("done.");
    }

    private static void startTimer() {
        Timer t = new Timer();
        t.schedule(new TimerTask(){
            public void run() { log.info("Timer fired."); }
        }, 1000*10);
    }
    
}
