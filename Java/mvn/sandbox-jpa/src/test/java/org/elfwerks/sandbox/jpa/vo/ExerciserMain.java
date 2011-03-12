package org.elfwerks.sandbox.jpa.vo;

import java.util.List;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;
import javax.persistence.Query;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class ExerciserMain {
    private static final Log log = LogFactory.getLog(ExerciserMain.class);
    private static final String persistenceUnitName = "sandbox-jpa";
    private static EntityManagerFactory emf;
    
    
    @SuppressWarnings("unchecked")
    public static void main(String[] args) {
        emf = Persistence.createEntityManagerFactory(persistenceUnitName);
        EntityManager em = emf.createEntityManager();
        try {
            log.debug("Starting dump of Bags.");
            Query q = em.createQuery("FROM BoxVO");
            List<BoxVO> boxList = (List<BoxVO>)q.getResultList();
            for (BoxVO b: boxList) {
                System.out.println("Box: "+b.getName());
                for (BoxExtraVO e: b.getExtras()) {
                    System.out.println("      ["+e.getId()+"] "+e.getName()+" : "+e.getValue());
                    
                    @SuppressWarnings("unused")
                    String ident = e.getIdentity();
                }
            }
        
        
        } catch (Throwable ex) {
            log.fatal(ex);
            //throw ex;
        } finally {
            em.close();
        }
        log.info("Done.");
    }

}
