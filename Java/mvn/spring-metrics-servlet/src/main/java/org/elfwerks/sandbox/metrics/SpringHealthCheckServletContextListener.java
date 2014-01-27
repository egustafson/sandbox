package org.elfwerks.sandbox.metrics;

import javax.servlet.ServletContext;
import javax.servlet.ServletContextEvent;
import javax.servlet.annotation.WebListener;

import org.springframework.web.context.WebApplicationContext;
import org.springframework.web.context.support.WebApplicationContextUtils;

import com.codahale.metrics.health.HealthCheckRegistry;
import com.codahale.metrics.servlets.HealthCheckServlet.ContextListener;

@WebListener
public class SpringHealthCheckServletContextListener extends ContextListener {
    
    protected WebApplicationContext springContext;

    @Override
    public void contextInitialized(ServletContextEvent event) {
        ServletContext sc = event.getServletContext();
        springContext = WebApplicationContextUtils.getWebApplicationContext(sc);
        super.contextInitialized(event);
    }
    
    protected HealthCheckRegistry getHealthCheckRegistry() {
        return springContext.getBean(HealthCheckRegistry.class);
    }

}
