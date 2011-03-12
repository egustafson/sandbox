package org.elfwerks.sandbox.json;


import static org.junit.Assert.*;

import org.json.simple.JSONObject;
import org.junit.Test;

public class TestEncoding {

	@Test
	public void testBasicEncoding() {
		JSONObject obj = new JSONObject();
		assertEquals(obj.toJSONString(),"{}");
	}
}
