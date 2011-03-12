package org.elfwerks.sandbox.env;

import java.util.Map;
import java.util.Set;

public class Dump {

	public static void main(String[] args) {
		Set<Map.Entry<Object, Object>> sysProps = System.getProperties().entrySet(); 
		System.out.println(" ----- System Properties ----- ");
		for (Map.Entry<Object, Object> kv : sysProps) {
			String key = (String)kv.getKey();
			String val = (String)kv.getValue();
			System.out.println(key + " = "+val);
		}
		Map<String, String> sysEnv = System.getenv();
		System.out.println(" ----- Environment Variables ----- ");
		for (Map.Entry<String, String> env : sysEnv.entrySet()) {
			System.out.println(env.getKey() + " = " + env.getValue());
		}
		
	}

}
