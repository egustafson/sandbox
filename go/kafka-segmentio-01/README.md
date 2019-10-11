Kafka Go Producer/Consumer using segmentio/kafka-go
===================================================

Superceeds sandox example 'kafka-sarama'.  It appears, at the time of
this writing, that https://github.com/segmentio/kafka-go is the new
darling.

Uses:
: https://github.com/segmentio/kafka-go

Prerequisites
-------------
1. Kafka running -- see:
   https://docs.confluent.io/current/quickstart/ce-docker-quickstart.html

2. Place an alias in /etc/hosts for 'broker' -> localhost.  This is an
   artifact of how kafka is run.  The broker knows itself as 'broker',
   but that host name is obscured by docker's network.  However the
   port is mapped to localhost.  (Read the docker-compose script.)

To-Do:  Implement example a docker containers and have the `Makefile`
run the containers in the proper docker network.



References
----------

* https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26
