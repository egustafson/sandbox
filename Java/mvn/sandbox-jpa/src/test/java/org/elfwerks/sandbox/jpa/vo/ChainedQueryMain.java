package org.elfwerks.sandbox.jpa.vo;

import java.util.List;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;
import javax.persistence.Query;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

/** This program is intended to show how to "fault in"
 * multiple JPA Entity objects using a single JPA Query.
 * The goal is to reduce queries to the DB when navigating
 * Entities by forming a query that loads all the objects
 * in one pass.
 */
public class ChainedQueryMain {
	private static final Log log = LogFactory.getLog(ChainedQueryMain.class);
	private static final String persistenceUnitName = "sandbox-jpa";
	private static EntityManagerFactory emf;

	public static void main(String[] args) {
        emf = Persistence.createEntityManagerFactory(persistenceUnitName);
        try {
        	//navigation();
        	//faultIn();
        	exhaustEMs();
        } catch (Throwable ex) {
            log.fatal(ex);
            //throw ex;
        }
        log.info("Done.");
	}
	
	@SuppressWarnings({ "unchecked", "unused" })
	private static void navigation() throws Exception {
        EntityManager em = emf.createEntityManager();
        try {
            log.debug("Starting dump of Bags.");
            Query q = em.createQuery("FROM BoxVO");
            List<BoxVO> boxList = (List<BoxVO>)q.getResultList();
            for (BoxVO b: boxList) {
                System.out.println("Box: "+b.getName());
                for (BoxExtraVO e: b.getExtras()) {
                    System.out.println("      ["+e.getId()+"] "+e.getName()+" : "+e.getValue());
                    String ident = e.getIdentity();
                }
            }
        } finally {
        	em.close();
        }
	}
	
	@SuppressWarnings({ "unchecked", "unused" })
	private static void faultIn() throws Exception {
		EntityManager em = emf.createEntityManager();
		try {
			log.debug("Starting 'fault in' dump of Bags.");
			Query q = em.createQuery("SELECT b FROM BoxVO b");
			Object result = q.getResultList();
			List<BoxVO> boxList = (List<BoxVO>) result;
			
			Query q2 = em.createQuery("SELECT be FROM BoxExtraVO be");
			result = q2.getResultList();
			
			for (BoxVO b: boxList) {
                System.out.println("Box: "+b.getName());
                for (BoxExtraVO e: b.getExtras()) {
                    System.out.println("      ["+e.getId()+"] "+e.getName()+" : "+e.getValue());
                    
                    String ident = e.getIdentity();
                }
			}
		} finally {
			em.close();
		}
	}
	
	@SuppressWarnings("unused")
	private static void exhaustEMs() throws Exception {
		for (int ii = 0; ii < 1000000; ii++) {
			EntityManager em = emf.createEntityManager();
			em.getTransaction().begin();
			em.getTransaction().rollback();
			em.close();
			//Thread.sleep(1);
		}
	}
	
}
