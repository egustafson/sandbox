package org.elfwerks;

public class PersonHandler {

	private Person person;
	
	public void setPerson(Person p) {
		person = p;
	}
	
	public String savePerson() {
		person.save();
		return "success";
	}
}
