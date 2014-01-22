package org.elfwerks.sandbox.jersey;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import org.springframework.stereotype.Component;

@Component
@Path("/")
@Produces(MediaType.APPLICATION_JSON)
public class ExampleResource {

  @GET
  public ExampleRepresentation doGet() {
    return new ExampleRepresentation();
  }
  
}
