package org.elfwerks.sandbox.jetty.spring;

import java.io.IOException;

import javax.servlet.http.HttpServletResponse;

import org.joda.time.DateTime;
import org.joda.time.format.DateTimeFormatter;
import org.joda.time.format.ISODateTimeFormat;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;

@Controller
public class SandboxController {

  private final static DateTimeFormatter isoFmt = ISODateTimeFormat.dateTime();
  
  @RequestMapping(method=RequestMethod.GET, value="/*")
  public void welcome(HttpServletResponse response) throws IOException {
    DateTime now = new DateTime();
    response.setStatus(200);
    response.getWriter().print("Welcome, the time is: " + isoFmt.print(now));
  }
  
}
