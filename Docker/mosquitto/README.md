Mosquitto Broker
================

* Start with `docker-compose`

CLI Client
----------

* `apt install mosquitto-clients`

Confirm Working
---------------

1. `mosqitto_sub -t test/topic`
2. `mosquitto_pub -t test/topic "test message 1"`

