package org.elfwerks.sandbox.spel;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class DemoSpringBean {

    private String name;
    private String version;
    private int port;
    
    @Value("#{appConfig[name]}")
    public void setName(String n) { name = n; }
    public String getName() { return name; }
    
    @Value("#{appConfig[version]}")
    public void setVersion(String v) { version = v; }
    public String getVersion() { return version; }
    
    @Value("#{appConfig[port]}")
    public void setPort(int p) { port = p; }
    public int getPort() { return port; }
    
    public String toString() {
        StringBuilder sb = new StringBuilder(name);
        sb.append('(').append(version).append(')')
        .append(" port: ").append(port);
        return sb.toString();
        
    }
}
