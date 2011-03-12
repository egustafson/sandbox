package org.elfwerks.sandbox.springjdbc;

import java.util.Collection;

/*  Table (id INTEGER, name VARCHAR, value VARCHAR)
 */

public interface NameValueTableDAO {

    public Collection<NameValuePair> dump();
    public void add(String name, String value);
    public void remove(String name);
    public Collection<NameValuePair> lookup(String name);
    public NameValuePair byId(int id);

    public void utilInitSchema();
    
}
