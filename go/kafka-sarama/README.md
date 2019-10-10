Kafka Go Producer/Consumer
==========================

Note:  _SUPERSEDED_ by the segmentio/kafka-go implementation.

Uses
: https://github.com/Shopify/sarama
: https://github.com/wvanbergen/kafka  (likely obsolete)

Prerequisites
-------------
1. Kafka running -- see `docker-start-kafka` directory for dev-docker
   startup.

2. Place an alias in /etc/hosts for 'broker' -> localhost.  This is an
   artifact of how kafka is run.  (The broker knows itself as
   'broker', but port mapped through docker to listen on localhost.)

References
----------

* https://medium.com/rahasak/kafka-consumer-with-golang-a93db6131ac2
* https://medium.com/rahasak/kafka-producer-with-golang-fab7348a5f9a
