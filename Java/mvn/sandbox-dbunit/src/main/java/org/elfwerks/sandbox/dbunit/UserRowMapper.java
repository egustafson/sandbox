package org.elfwerks.sandbox.dbunit;

import java.sql.ResultSet;
import java.sql.SQLException;

import org.springframework.jdbc.core.RowMapper;

public class UserRowMapper implements RowMapper {
	
	public User mapRow(ResultSet rs, int rowNum) throws SQLException {
		return new User(rs.getString("username"), rs.getString("email"));
	}
}