#
#
## Imports

from mmaplog import BaseLogBroker, BaseConsumer, Producer


## ##################################################

def test_function_null():
    assert True

def test_create_baselog():
    bl = BaseLogBroker()
    assert isinstance(bl.put(None), (int))
    assert bl.get_consumer != None
    assert bl.get_producer != None


def test_base_consumer():
    bl = BaseLogBroker()
    consumer = bl.get_consumer()
    assert consumer.next() == None
    assert consumer.next(100) == None


def test_producer():
    bl = BaseLogBroker()
    producer = bl.get_producer()
    assert isinstance(producer.put(None), (int))
    assert isinstance(producer.put("value"), (int))


## Local Variables:
## mode: python
## End:
