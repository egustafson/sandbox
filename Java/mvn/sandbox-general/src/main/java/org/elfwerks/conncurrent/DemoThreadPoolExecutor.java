package org.elfwerks.conncurrent;

import java.util.concurrent.Callable;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.FutureTask;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;


public class DemoThreadPoolExecutor {

    static class Task implements Callable<Integer> {

        private int count;
        
        public Task(int c) {
            count = c;
        }
        
        @Override
        public Integer call() throws Exception {
            // Work complete - no-op
            System.out.println(count);
            return new Integer(1);
        }
    }
    
    /**
     * @param args
     */
    public static void main(String[] args) {
        ThreadPoolExecutor ex = new ThreadPoolExecutor(30, 30, 30, TimeUnit.SECONDS, new LinkedBlockingQueue<Runnable>());
        for (int ii = 0; ii < 10000; ii++) {
            FutureTask<Integer> task = new FutureTask<Integer>(new Task(ii));
            ex.execute(task);
        }
        ex.shutdown();
        try {
            int activeTasks = ex.getActiveCount();
            System.err.print("Shutdown started - "+activeTasks+" tasks active.");
            Thread.sleep(300);
            System.err.print("After shutdown delay.");
            if ( !ex.isShutdown() || !ex.isTerminated() ) {
                Thread.sleep(1000);
            }
        } catch (InterruptedException ex1) {
            // TODO Auto-generated catch block
            ex1.printStackTrace();
        }
        if ( !ex.isShutdown() || !ex.isTerminated() ) {
            ex.shutdownNow();
        }
        System.err.println("Done. ----------");
    }

}
