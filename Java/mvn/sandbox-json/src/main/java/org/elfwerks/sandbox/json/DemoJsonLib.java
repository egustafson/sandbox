package org.elfwerks.sandbox.json;

import java.util.LinkedList;
import java.util.List;

import net.sf.json.JSONObject;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class DemoJsonLib {
	private static final Log log = LogFactory.getLog(DemoJsonLib.class);
	
    public static void main( String[] args ) {
    	log.info("Using JSON-Lib");
    	log.info("--- Simple Bean Example ---");
    	exerciseBean(createSimpleBean());
    	log.info("--- Complex Bean Example ---");
    	//
    	// Unmarshaling the complex bean fails as JSON-Lib can't
    	// reach into the List<> class.
    	//
    	exerciseBean(createComplexBean());
    }
    
    public static void exerciseBean(Object bean) {
    	log.info("The Bean: "+bean.toString());
    	JSONObject jobj = JSONObject.fromObject(bean);
    	String jsonText = jobj.toString();
    	log.info("JSON:  "+jobj);
    	
    	JSONObject jobj2 = JSONObject.fromObject(jsonText);
    	Object bean2 = JSONObject.toBean(jobj2, bean.getClass());

    	log.info("Reconstituted bean is a: "+bean2.getClass().getCanonicalName());
    	log.info("Bean': "+bean2.toString());
    }
    
    public static ComplexMessageBean createComplexBean() {
    	List<MessageBean> msgList = new LinkedList<MessageBean>();
    	for(int ii = 1; ii < 3; ii++) {
    		MessageBean msg = createSimpleBean();
    		msg.setCount(ii);
    		msgList.add(msg);
    	}
    	ComplexMessageBean cBean = new ComplexMessageBean();
    	cBean.setPrefix("complex-bean");
    	cBean.setMessages(msgList);
    	return cBean;
    }
    
    
    public static MessageBean createSimpleBean() {
    	MessageBean msg = new MessageBean();
    	msg.setTitle("MsgTitle");
    	msg.setSequence(21);
    	msg.setMessage("The message.");
    	msg.setCount(1);
    	return msg;
    }
}
