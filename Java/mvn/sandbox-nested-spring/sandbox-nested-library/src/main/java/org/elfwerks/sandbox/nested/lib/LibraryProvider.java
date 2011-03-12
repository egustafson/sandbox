package org.elfwerks.sandbox.nested.lib;

import org.elfwerks.sandbox.nested.api.Factory;
import org.elfwerks.sandbox.nested.api.Provider;

public class LibraryProvider implements Provider {

    @Override
    public String getName() {
        return "LibraryProvider";
    }

    @Override
    public void setFactory(Factory f) {
        f.register(this);
    }

    
}
