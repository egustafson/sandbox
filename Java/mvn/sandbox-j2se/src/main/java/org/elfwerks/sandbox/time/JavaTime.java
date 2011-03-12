package org.elfwerks.sandbox.time;

public class JavaTime {

    public static void main(String[] args) {
        long now = System.currentTimeMillis();
        long secondsFromEpoc = now/1000;

        System.out.println("Seconds since Java Epoc:  "+secondsFromEpoc);
    }

}
