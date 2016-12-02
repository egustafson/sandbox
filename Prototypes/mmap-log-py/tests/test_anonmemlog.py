#
#
## Imports

from mmaplog import AnonMemLogBroker


## ##################################################

def test_create_log():
    lg = AnonMemLogBroker()
    assert lg != None

def test_anonmem_put():
    lg = AnonMemLogBroker()
    assert lg.put("x") == 0
    assert lg.put("x") == 1
    assert lg.put("x") == 2

def test_anonmem_produce():
    lg = AnonMemLogBroker()
    pr = lg.get_producer()
    assert pr.put("x") == 0
    assert pr.put("x") == 1
    assert pr.put("x") == 2
    for x in range(1, 1000):
        pr.put("x")

def test_anonmem_consume_one():
    src = "x_test_value_x"
    lg = AnonMemLogBroker()
    lg.put(src)
    consumer = lg.get_consumer()
    dst = consumer.next()
    assert src == dst
