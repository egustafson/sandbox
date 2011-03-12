package org.elfwerks.sandbox.dbunit;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name="roles")
public class Role {

	private int id;
	private String rolename;

	@Id
	@GeneratedValue(strategy=GenerationType.AUTO)
	public int getId() { return id; }
	protected void setId(int id) { this.id = id; } 
	
	@Column(name="rolename", unique=true, nullable=false, length=32)
	public String getRolename() { return rolename; }
	public void setRolename(String rolename) { this.rolename = rolename; }
	
}
