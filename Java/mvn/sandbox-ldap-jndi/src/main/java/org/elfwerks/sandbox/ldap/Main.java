package org.elfwerks.sandbox.ldap;

import java.util.Properties;

import javax.naming.Context;
import javax.naming.NameClassPair;
import javax.naming.NamingEnumeration;
import javax.naming.NamingException;
import javax.naming.directory.Attribute;
import javax.naming.directory.Attributes;
import javax.naming.directory.DirContext;
import javax.naming.directory.InitialDirContext;

public class Main 
{
    public static void main( String[] args ) throws Exception {
    	Properties env = new Properties();
    	env.put(Context.INITIAL_CONTEXT_FACTORY,"com.sun.jndi.ldap.LdapCtxFactory");
    	env.put(Context.PROVIDER_URL, "ldap://192.168.184.128:389/");
    	env.put(Context.SECURITY_AUTHENTICATION,"simple");
    	env.put(Context.SECURITY_PRINCIPAL,"cn=Eric Gustafson,ou=People,dc=example,dc=com"); // specify the username
    	env.put(Context.SECURITY_CREDENTIALS,"password");           // specify the password
    	try {
    		DirContext rootCtx = new InitialDirContext(env);
    		NamingEnumeration<NameClassPair> e = rootCtx.list("dc=example,dc=com");
    		while ( e.hasMore() ) {
    			NameClassPair ncPair = e.next();
    			System.out.println("  "+ncPair.getName()+": "+ncPair.getClassName());
    		}
    		System.out.println("----------");
    		Attributes exAttrs = rootCtx.getAttributes("dc=example,dc=com");
    		NamingEnumeration<String> attrNames = exAttrs.getIDs();
    		while (attrNames.hasMore()) {
    			String name = attrNames.next();
    			Attribute a = exAttrs.get(name);
    			NamingEnumeration<?> values = a.getAll();
    			while (values.hasMore()) {
    				Object v = values.next();
    				System.out.println("  "+name+": "+v.toString());
    			}
    		}
    		
    		Attributes attrs = rootCtx.getAttributes("cn=Eric Gustafson,ou=People,dc=example,dc=com");
    		System.out.println("Last Name: " + attrs.get("sn").get());
    		rootCtx.close();
    	} catch (NamingException e) {
    		System.err.println("Problem getting attribute: " + e);
    	}    	
    	System.out.println("Done.");
    }
}
