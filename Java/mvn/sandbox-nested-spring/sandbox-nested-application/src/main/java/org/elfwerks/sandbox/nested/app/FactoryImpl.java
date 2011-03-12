package org.elfwerks.sandbox.nested.app;

import java.net.URL;
import java.util.Enumeration;
import javax.annotation.PostConstruct;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.xml.XmlBeanDefinitionReader;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.context.support.GenericApplicationContext;
import org.springframework.core.io.UrlResource;

import org.elfwerks.sandbox.nested.api.Factory;
import org.elfwerks.sandbox.nested.api.Provider;

public class FactoryImpl implements Factory, ApplicationContextAware {
    private static final Log log = LogFactory.getLog(FactoryImpl.class);
    private static final String PROVIDER_CTX_FILENAME = "providerContext.xml";
    
    private ApplicationContext applicationContext;
    
    @Override
    public void setApplicationContext(ApplicationContext applicationContext) throws BeansException {
        this.applicationContext = applicationContext;
    }

    @PostConstruct
    public void initialize() {
        log.info("Factory initialization.");
        try {
            Enumeration<URL> providerCtx = getClass().getClassLoader().getResources(PROVIDER_CTX_FILENAME);
            while (providerCtx.hasMoreElements()) {
                URL ctxUrl = providerCtx.nextElement();
                log.info("Found provider context: "+ctxUrl.toString());
                
                GenericApplicationContext ctx = new GenericApplicationContext(applicationContext);
                XmlBeanDefinitionReader xmlReader = new XmlBeanDefinitionReader(ctx);
                xmlReader.loadBeanDefinitions(new UrlResource(ctxUrl));
                ctx.refresh();
            }
        } catch (Exception ex) {
            log.error("Caught "+ex.getClass().getName()+" - (reason:"+ex.getLocalizedMessage()+")", ex);
        }
        log.info("Factory initialization finished.");
    }
    
    public void register(Provider p) {
        log.info("Provider: "+p.getName()+" registered.");
    }
    
}
