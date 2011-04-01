package org.elfwerks.sandbox.threadpool;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.core.task.TaskExecutor;

public class PooledWorker {
    private final Log log = LogFactory.getLog(this.getClass());
    
    private class WorkTask implements Runnable {
        private String message;
        
        public WorkTask(String m) {
            message = m;
        }
        
        public void run() {
            try {
                Thread.sleep(1000); /* sleep 1s */
            } catch (InterruptedException ex) {
                log.warn(ex);
            } 
            log.info("Processed message: ["+message+"]");
        }
        
    }

    private TaskExecutor taskExecutor;
    
    public PooledWorker() { /* null constructor */ }
    
    public void setTaskExecutor(TaskExecutor te) {
        taskExecutor = te;
    }
    
    public void createWork() {
        log.debug("Begin - creating messages.");
        for (int ii = 0; ii < 15; ii++) {
            String msg = "Message-"+ii;
            taskExecutor.execute(new WorkTask(msg));
            log.info("enqueued: ["+msg+"]");
        }
        log.debug("Finished - creating messages.");
    }

}
