package org.elfwerks.sandbox.dbunit;

import java.util.List;

import javax.sql.DataSource;

public interface UserDao {
	
	public void setDataSource(DataSource dataSource);
	public List<User> getAllUsers();

}
