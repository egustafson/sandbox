package org.elfwerks.sandbox.pool;

import java.io.IOException;
import java.util.Properties;

import javax.annotation.PostConstruct;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Required;
import org.springframework.context.ApplicationContext;
import org.springframework.core.io.Resource;
import org.springframework.util.AntPathMatcher;

public class PoolLoader {
	private final Log log = LogFactory.getLog(this.getClass());
	
	@Autowired
	private PoolManager poolManager;
	@Autowired
	private ApplicationContext appCtx;
	
	private String resourcePath;
	@Required
	public void setResourcePath(String r) { resourcePath = r; }

	@PostConstruct
	public void init() throws IOException {
		log.info("Performing initialization with ["+resourcePath+"]");
		Resource[] resources = appCtx.getResources(resourcePath);
		AntPathMatcher pm = new AntPathMatcher();
		for (Resource r : resources) {
			log.debug("Resource: "+r.getURL()+" ["+r.getURI()+"]");
			Properties props = new Properties();
			try {
				props.load(r.getInputStream());
				String name = pm.extractPathWithinPattern(resourcePath, r.getFilename());
				PoolElement pe = appCtx.getBean(PoolElement.class);
				pe.configure(name, props);
				poolManager.register(pe);
			} catch (IOException ex) {
				log.warn("Failed to create the PoolElement["+r.getFilename()+"]");
			} catch (PoolElConfigurationException ex) {
				log.warn("Failed to CONFIGURE PoolElement["+r.getFilename()+"]");
			}
		}
		
		
		poolManager.register(null);
		log.info("initialization complete.");
	}

	
}
