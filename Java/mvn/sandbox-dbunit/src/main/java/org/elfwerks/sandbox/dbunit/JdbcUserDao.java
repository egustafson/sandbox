package org.elfwerks.sandbox.dbunit;

import java.util.List;

import javax.sql.DataSource;

import org.springframework.jdbc.core.JdbcTemplate;

public class JdbcUserDao implements UserDao {
	private JdbcTemplate jdbcTemplate;

	public void setDataSource(DataSource dataSource) {
		jdbcTemplate = new JdbcTemplate(dataSource);
	}
	
	@SuppressWarnings("unchecked")
	public List<User> getAllUsers() {
		final String sql = "SELECT username, email FROM USERS";
		return jdbcTemplate.query(sql, new UserRowMapper());
	}
}
