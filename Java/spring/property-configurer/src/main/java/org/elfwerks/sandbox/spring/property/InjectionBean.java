package org.elfwerks.sandbox.spring.property;

public class InjectionBean {
	String  name = "nil";
	Integer ver  = 0;
	
	public void setName(String n) { name = n; }
	public void setVer(Integer v) { ver = v; }
	
	public String toString() {
		return "Name: ["+name+"], ver("+ver+")";
	}
}
