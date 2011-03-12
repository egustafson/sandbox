package org.elfwerks.sandbox.nested.app;

import org.elfwerks.sandbox.nested.api.Factory;
import org.elfwerks.sandbox.nested.api.Provider;

public class CoreProvider implements Provider {

    private String name = "CoreProvider";
    
    @Override
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
    
    @Override
    public void setFactory(Factory f) {
        f.register(this);
    }

}
