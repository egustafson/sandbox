package org.elfwerks.sandbox.bytebuffer;

import java.nio.ByteBuffer;

public class ExerciseSlice {

    /**
     * @param args
     */
    public static void main(String[] args) {
        ByteBuffer buf = ByteBuffer.allocate(1024);
        buf.position(64);
        buf.limit(128);
        
        buf.mark();
        for (byte ii = 0; ii < 64; ii++) {
            buf.put(ii);
        }
        buf.reset();
        
        ByteBuffer slice = buf.slice();
        byte[] bslice = new byte[slice.capacity()];
        slice.get(bslice);
        
        System.out.println("Final byte array is "+bslice.length+" bytes in length.");

    }

}
