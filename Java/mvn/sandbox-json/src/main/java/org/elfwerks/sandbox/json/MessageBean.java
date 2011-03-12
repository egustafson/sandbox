package org.elfwerks.sandbox.json;

public class MessageBean {

	private String title;
	private Integer sequence;
	private String message;
	private int count;
	
	public String getTitle() {
		return title;
	}
	public void setTitle(String title) {
		this.title = title;
	}
	public Integer getSequence() {
		return sequence;
	}
	public void setSequence(Integer sequence) {
		this.sequence = sequence;
	}
	public String getMessage() {
		return message;
	}
	public void setMessage(String message) {
		this.message = message;
	}
	public int getCount() {
		return count;
	}
	public void setCount(int count) {
		this.count = count;
	}

	@Override
	public String toString() {
		return "Message["+title+" - seq("+sequence+", "+count+"): "+message+"]";
	}
}
