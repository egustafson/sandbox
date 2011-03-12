package org.elfwerks.sandbox.springjdbc;

public class NameValuePair {

    private int id;
    private String name;
    private String value;

    public NameValuePair() { /* do nothing */ }

    public NameValuePair(int i, String n, String v) {
        id = i;
        name = n;
        value = v;
    }

    public int getId() { return id; }
    public void setId(int i) { id = i; }

    public String getName() { return name; }
    public void setName(String n) { name = n; }

    public String getValue() { return value; }
    public void setValue(String v) { value = v; }

    public String toString() {
        return "{(" + id + ") " + name + ": " + value + "}";
    }

}
