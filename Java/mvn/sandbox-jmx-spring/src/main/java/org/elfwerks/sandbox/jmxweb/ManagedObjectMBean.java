package org.elfwerks.sandbox.jmxweb;

public interface ManagedObjectMBean {

    public enum State {ACTIVE, PASSIVE};
    
    public Integer getValue();
    public void setValue(Integer value);
    public Integer getCounter();
    public String getState();
    
    public void activate();
    public void pacify();
    
    
}
