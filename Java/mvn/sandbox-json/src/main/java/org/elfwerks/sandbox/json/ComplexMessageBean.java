package org.elfwerks.sandbox.json;

import java.util.List;

public class ComplexMessageBean {

	String prefix;
	List<MessageBean> messages;
	
	
	public String getPrefix() {
		return prefix;
	}
	public void setPrefix(String prefix) {
		this.prefix = prefix;
	}
	public List<MessageBean> getMessages() {
		return messages;
	}
	public void setMessages(List<MessageBean> messages) {
		this.messages = messages;
	}
	
	@Override
	public String toString() {
		StringBuilder sb = new StringBuilder("ComplexMessage("+prefix+": "+messages.size()+" messages)");
		for(MessageBean m : messages) {
			sb.append("\n  "+m.toString());
		}
		return sb.toString();
	}
	
}
