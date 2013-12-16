package org.elfwerks.sandbox.dropwizard;

import org.springframework.context.ApplicationContext;
import org.springframework.context.support.ClassPathXmlApplicationContext;

import com.yammer.dropwizard.Service;
import com.yammer.dropwizard.config.Bootstrap;
import com.yammer.dropwizard.config.Environment;

public class HelloWorldService extends Service<HelloWorldConfiguration> {

  @Override
  public void initialize(Bootstrap<HelloWorldConfiguration> bootstrap) {
    bootstrap.setName("hello-world");
  }
  
  @Override
  public void run(HelloWorldConfiguration configuration, Environment environment) {
    final String template = configuration.getTemplate();
    final String defaultName = configuration.getDefaultName();
    environment.addResource(new HelloWorldResource(template, defaultName));
    environment.addHealthCheck(new TemplateHealthCheck(template));
  }
  
  public static void main(String[] args) throws Exception {
    ApplicationContext context = new ClassPathXmlApplicationContext("hello-spring-context.xml");
    new HelloWorldService().run(args);
  }

}
