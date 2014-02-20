package org.elfwerks.sandbox.thymeleaf;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
public class ExampleController {

    @RequestMapping("/thymeleaf")
    public String response() {
        return "example-view";
    }
    
}
