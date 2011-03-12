package org.elfwerks.sandbox.corba;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

import org.elfwerks.sandbox.servant.PingServicePOA;

public class PingServant extends PingServicePOA {
    private static final Log log = LogFactory.getLog(PingServant.class);

    @Override
    public void ping() {
        log.info("ping invoked.");
    }

}
