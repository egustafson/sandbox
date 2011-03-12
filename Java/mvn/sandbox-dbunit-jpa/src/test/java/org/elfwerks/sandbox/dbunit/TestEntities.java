package org.elfwerks.sandbox.dbunit;

import java.util.Collection;

import javax.annotation.Resource;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import javax.sql.DataSource;

import static org.junit.Assert.*;

import org.dbunit.dataset.IDataSet;
import org.dbunit.dataset.xml.FlatXmlDataSet;
import org.elfwerks.unittest.AbstractJunit4DatabaseTestCase;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.test.context.ContextConfiguration;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.transaction.annotation.Transactional;

@RunWith(SpringJUnit4ClassRunner.class)
@ContextConfiguration(locations="PersistenceTests-context.xml")
public class TestEntities extends AbstractJunit4DatabaseTestCase {

	@Resource(name="dataSource")
	DataSource dataSource;
	
	@PersistenceContext
	EntityManager em;
	
	@Override
	protected DataSource getDataSource() {
		return dataSource;
	}

	@Override
	protected IDataSet getDataSet() throws Exception {
		return new FlatXmlDataSet(getClass().getResource("TestData.xml"));
	}

	@Test
	public void testGetRolename() {
		Role r = em.find(Role.class, 1);
		assertNotNull(r);
		if ( r != null ) {
			assertEquals(r.getRolename(), "test-role");
		}
	}
	
	@Test
	@Transactional
	public void testCreateRolename() {
		final String roleName = "test-inserted-role";
		Role role = new Role();
		role.setRolename(roleName);
		em.persist(role);
		int id = role.getId();
		role = em.find(Role.class, id);
		assertNotNull(role);
		assertEquals(role.getRolename(), roleName);
	}
	
	@Test
	public void testGetUsername() {
		User u = em.find(User.class, 1);
		assertNotNull(u);
		if ( u != null ) {
			assertEquals(u.getUsername(), "test-user");
			assertEquals(u.getRealname(), "Test User");
		}
	}
	
	@Test
	@Transactional
	public void testCreateUsername() {
		final String userName = "test-inserted-user";
		final String realName = "Inserted User";
		User user = new User();
		user.setUsername(userName);
		user.setRealname(realName);
		em.persist(user);
		int id = user.getId();
		user = em.find(User.class, id);
		assertNotNull(user);
		assertEquals(user.getUsername(), userName);
		assertEquals(user.getRealname(), realName);
	}
	
	@Test
	@Transactional
	public void testUserRoleRelationship() {
		User u = em.find(User.class, 1);
		Collection<Role> roles = u.getRoles();
		assertEquals(roles.size(), 1);
		Role r = roles.iterator().next();
		assertEquals(r.getRolename(), "test-role");
	}

	@Test
	@Transactional
	public void testUserRoleCreation() {
		User u = em.find(User.class, 2);
		Role r = em.find(Role.class, 2);
		String rolename = r.getRolename();
	
		u.getRoles().clear();
		u.getRoles().add(r);
		em.flush();
		
		u = em.find(User.class, 2);
		r = u.getRoles().iterator().next();
		
		assertEquals(r.getRolename(), rolename);
	}
	
}
