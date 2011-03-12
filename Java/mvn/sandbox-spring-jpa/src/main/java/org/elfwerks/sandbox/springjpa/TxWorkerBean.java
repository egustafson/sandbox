package org.elfwerks.sandbox.springjpa;

import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

@Transactional(propagation = Propagation.SUPPORTS)
public class TxWorkerBean {
	private static final Log log = LogFactory.getLog(TxWorkerBean.class);

    @PersistenceContext(unitName="sandbox-persistence")
    private EntityManager em;
    
    @PersistenceContext(unitName="sandbox-second-persistence")
    private EntityManager emSecond;
    
    public void lookupNoTx() {
    	log.debug("in lookupNoTx() method");
    	em.clear();
    	emSecond.clear();
    }
    
    @Transactional(propagation = Propagation.REQUIRED)
    public void modifyRequireTx() {
    	log.debug("in modifyRequireTx() method");
    	em.clear();
    }
    
    @Transactional(propagation = Propagation.REQUIRED, timeout=30) // timeout doesn't work w/ JpaTx
    public void nestedRequireTx() {
    	log.debug("in nestedRequireTx - before nested call");
    	this.lookupNoTx();
    	log.debug("in nestedRequireTx - afer nested call");
    }
	
	
}
