#
#
## Imports


## ##################################################
##
class BaseLogBroker(object):

    def __init__(self):
        pass

    def get_consumer(self):
        return BaseConsumer(self)

    def get_producer(self):
        return Producer(self)

    def put(self, value):
        return 1  ## sequence number


## ##################################################
##
class BaseConsumer(object):

    def __init__(self, broker):
        self._broker = broker

    def seek(self, seqnum):
        raise NotImplementedError("BaseConsumer.seek() - not implemented")

    def next(self, timeout=None):
        return None


## ##################################################
##
class Producer(object):

    def __init__(self, broker):
        self._broker = broker

    def put(self, value):
        return self._broker.put(value)


## ############################################################
##
##   AnonMem Log -- based on anonymous memory mmap
##

from binascii import crc32
from mmap import mmap
from os import statvfs
from struct import pack, unpack, calcsize


AM_DEFAULT_MAX_LOG_SIZE = 1024 * 1024  ## 1 MB
#AM_INITIAL_SIZE = statvfs('/').f_bsize * 10
#AM_GROW_SIZE = statvfs('/').f_bsize

AM_HDR_STRUCT  = '=iHII'  # crc32, ver, seq, size
AM_HDR_SIZE    = calcsize(AM_HDR_STRUCT)
AM_HDR_VERSION = 0

## ##################################################
##
class AnonMemLogBroker(BaseLogBroker):

    def __init__(self, size=AM_DEFAULT_MAX_LOG_SIZE):
        self.next_seq = 0
        self._mm = mmap(-1, size)
        self._mm.seek(0)

    def put(self, value):
        vlen = len(value)
        end = self._mm.tell() + AM_HDR_SIZE + vlen
        if end > self._mm.size():
            raise "Overflow"
        crc = crc32(value)
        seq = self.next_seq
        hdr = pack(AM_HDR_STRUCT, crc, AM_HDR_VERSION, seq, vlen)
        self.next_seq += 1
        self.write(hdr)
        self.write(value)
        return seq

    def get_consumer(self):
        return AnonMemConsumer(self)


## ##################################################
##
class AnonMemConsumer(BaseConsumer):

    def seek(self, seqnum):
        #
        # TODO
        #
        pass


    def next(self):
        #
        # TODO
        #
        pass




## Local Variables:
## mode: python
## End:
