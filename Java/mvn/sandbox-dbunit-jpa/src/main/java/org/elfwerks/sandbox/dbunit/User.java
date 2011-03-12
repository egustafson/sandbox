package org.elfwerks.sandbox.dbunit;

import java.util.Set;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

@Entity
@Table(name="users")
public class User {

	private int id;
	private String username;
	private String realname;
	private Set<Role> roles;
	
	@Id
	@GeneratedValue(strategy=GenerationType.AUTO)
	public int getId() { return id;	}
	protected void setId(int id) { this.id = id; }
	
	@Column(name="username", unique=true, nullable=false, length=64)
	public String getUsername() { return username;}
	public void setUsername(String username) { this.username = username; }
	
	@Column(name="real_name", length=255)
	public String getRealname() { return realname; }
	public void setRealname(String realname) { this.realname = realname; }
	
	@ManyToMany
	public Set<Role> getRoles() { return roles;	}
	public void setRoles(Set<Role> roles) { this.roles = roles;	}
	
}
