package org.elfwerks.sandbox.jmxweb;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.springframework.beans.factory.annotation.Required;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.mvc.AbstractController;

public class SimpleController extends AbstractController {

    private ManagedObject managedObject;

    @Required
    public void setManagedObject(ManagedObject mo) { managedObject = mo; }
    
    @Override
    protected ModelAndView handleRequestInternal(HttpServletRequest request, HttpServletResponse response) throws Exception {
        managedObject.incCounter();
        return new ModelAndView("index", "mo", managedObject);
    }

}
