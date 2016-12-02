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
        self.size = size
        self._mm = mmap(-1, size)
        self._mm.seek(0)

    def put(self, value):
        vlen = len(value)
        end = self._mm.tell() + AM_HDR_SIZE + vlen
        if end > self.size:
            raise "Overflow"
        crc = crc32(value)
        seq = self.next_seq
        hdr = pack(AM_HDR_STRUCT, crc, AM_HDR_VERSION, seq, vlen)
        self.next_seq += 1
        self._mm.write(hdr)
        self._mm.write(value)
        return seq

    def get_consumer(self):
        return AnonMemConsumer(self)


## ##################################################
##
class AnonMemConsumer(BaseConsumer):

    def __init__(self, broker):
        super(AnonMemConsumer, self).__init__(broker)
        self.pos = 0
        hdr = broker._mm[0:AM_HDR_SIZE]
        (crc, ver, seq, size) = unpack(AM_HDR_STRUCT, hdr)
        self.seq = seq

    def seek(self, seqnum):
        if self._broker.next_seq <= seqnum:
            raise "out of bounds"
        #
        # TODO
        #
        pass


    def next(self):
        if (self.pos + AM_HDR_SIZE) > self._broker.size:
            raise "out of bounds"
        hdr_end = self.pos + AM_HDR_SIZE
        (crc, ver, seq, size) = unpack(AM_HDR_STRUCT,
                                       self._broker._mm[self.pos:hdr_end])
        value_end = hdr_end + size
        if value_end > self._broker.size:
            raise "out of bounds"
        value = self._broker._mm[hdr_end:value_end]
        self.pos = value_end
        return value




## Local Variables:
## mode: python
## End:
