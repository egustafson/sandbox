package org.elfwerks.sandbox.json;

import java.util.ArrayList;
import java.util.List;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import com.google.gson.Gson;

public class DemoGson {
	private static final Log log = LogFactory.getLog(DemoGson.class);

	public static void main(String[] args) {
		log.info("--- Start - Using Gson ---");
    	log.info("--- Simple Bean Example ---");
    	exerciseBean(createSimpleBean());
    	log.info("--- Complex Bean Example ---");
    	//
    	// Note, GSON will create the List<MessageBean> from a LinkedList,
    	// but the creator (original bean) uses an ArrayList.  Both meet
    	// the contract defined in the definition of ComplexMessageBean, 
    	// which is simply a List<>.
    	exerciseBean(createComplexBean());
	}
	
    public static void exerciseBean(Object bean) {
    	log.info("The Bean: "+bean.toString());
    	Gson gson = new Gson();
    	String json = gson.toJson(bean);
    	log.info("JSON:  "+json);
    	
    	Object bean2 = gson.fromJson(json, bean.getClass());

    	log.info("Reconstituted bean is a: "+bean2.getClass().getCanonicalName());
    	log.info("Bean': "+bean2.toString());
    }
	
    public static ComplexMessageBean createComplexBean() {
    	List<MessageBean> msgList = new ArrayList<MessageBean>();
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
