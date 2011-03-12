package org.elfwerks.sandbox.jmxweb;

import org.springframework.jmx.export.annotation.ManagedResource;

public class ManagedObject implements ManagedObjectMBean {

    private Integer value = new Integer(0);
    private Integer counter = new Integer(0);
    private State state = State.ACTIVE;

    
    @Override
    public Integer getCounter() { return counter; }
    public void setCounter(Integer c) { counter = c; }
    public void incCounter() { counter += 1; }

    @Override
    public String getState() { return state.toString(); }

    @Override
    public Integer getValue() { return value; }
    @Override
    public void setValue(Integer v) { value = v; } 
    
    @Override
    public void activate() { state = State.ACTIVE; }
    @Override
    public void pacify() { state = State.PASSIVE; }
    
}
