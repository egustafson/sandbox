package org.elfwerks.sandbox.properties;

import java.io.InputStream;
import java.util.Properties;

public class DemoProperties {
	private static final String pFileName = "demoProperties.properties";
	
	public static void main(String[] args) throws Exception {
		InputStream is = DemoProperties.class.getClassLoader().getResourceAsStream(pFileName);
		Properties props = new Properties();
		props.load(is);
		props.list(System.out);
	}
}
