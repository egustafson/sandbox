package org.elfwerks.sandbox.spel;

import java.util.concurrent.ConcurrentHashMap;

import org.springframework.stereotype.Component;

@Component("appConfig")
public class AppConfig extends ConcurrentHashMap<String, String> {

    private static final long serialVersionUID = 6707459184549359668L;

    public AppConfig() {
        put("name", "Configuration Name");
        put("version", "1.2.0");
        put("port", "2120");
    }
    
}
