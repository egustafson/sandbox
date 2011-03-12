package org.elfwerks.sandbox.springjdbc;

import java.util.Collection;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

public class Main {
    private static final Log log = LogFactory.getLog(Main.class);
    private static final String appConfig = "applicationContext.xml";

    @SuppressWarnings("unused")
    public static void main(String[] args) {
        ApplicationContext ctx = new ClassPathXmlApplicationContext(appConfig);
        log.info("Loaded context: "+ctx.getDisplayName()+"("+ctx.getId()+")");
        
        
        NameValueTableDAO dao = ctx.getBean(NameValueTableDAO.class);
        dao.utilInitSchema();
        /* Exercise -- */

        dao.add("A", "Value-A");
        dao.add("B", "Value-B");
        dao.add("C", "Value-C");
        dao.add("D", "Value-D");
        
        NameValuePair nv = dao.byId(1);
        Collection<NameValuePair> result = dao.lookup("C");
        dao.remove("B");
        result = dao.dump();
        for (NameValuePair nvp : result) {
            log.debug(nvp);
        }
        
        
        /* end Exercise -- */
        System.out.println("Done.");
    }

}
