from binascii import crc32
from mmap import mmap, ACCESS_READ
from struct import pack, unpack, calcsize


## Notes:  Create a publishable (and subscribable) object that wraps an mmap.
##  * publish -> (binary)Length + string
##  * ? fixed length file & when pub crosses boundary, create next file?
##  ?? how to open a new mm file and grow it with out stream writing (ie zero size problem)
##  Create a reader example too (same example file, use click?

## Header format
##
##  <int>    :  crc32 of payload
##  <ushort> :  header version (currently 0)
##  <uint>   :  payload sequence number
##  <uint>   :  payload size in bytes


DEFAULT_SIZE = 1024 * 1024  # (1 Meg)

HDR_FMT  = '=iHII'  # crc32, ver, seq, size
HDR_SIZE = calcsize(HDR_FMT)


class JournalFullError(Exception):
    def __init__(self):
        None

    def __str__(self):
        return("JournalFullError")


class JournalConsumer(object):

    def __init__(self, filename):
        self._seq = 0
        self._eof = False
        self.filename = filename
        self.f = open(filename, "r")
        self.mm = mmap(self.f.fileno(), 0,
                       access=ACCESS_READ)

    def seek(self, index):
        None

    def seq(self):
        return self._seq

    def pos(self):
        return self.mm.tell()

    def size(self):
        return self.mm.size()

    def eof(self):
        return self._eof


    def next(self):
        if self._eof:
            raise StopIteration
        hdr = self.mm.read(HDR_SIZE)
        (crc, ver, seq, size) = unpack(HDR_FMT, hdr)
        self._seq = seq
        blob = self.mm.read(size)
        self._eof = (self.mm.size() - self.mm.tell()) <= 0
        return ( seq, blob )



class MMJournal(object):

    def __init__(self, filename, size=DEFAULT_SIZE):
        self.filename = filename
        self.seqnum = 0
        self.f = open(filename, "w+b")
        self.f.write('\0')
        self.f.flush()
        self.mm = mmap(self.f.fileno(), 0)
        self.mm.resize(size)
        self._size = size
        self.mm.seek(0)


    def available(self):
        if self.mm == None:
            return 0;
        return self.mm.size() - self.mm.tell()

    def size(self):
        return self._size


    def entries(self):
        return self.seqnum


    def close(self):
        self.mm.flush()
        self.mm.close()
        self.mm = None
        self.f.close()
        self.f = None


    def append(self, blob):
        if self.mm == None:
            raise JournalFullError()
        bl = len(blob)
        crc = crc32(blob)
        end = self.mm.tell() + HDR_SIZE + bl
        if end > self.mm.size():
            close_file = True
            self.mm.resize(end)
            self._size = end
        ver = 0
        hdr = pack(HDR_FMT, crc, ver, self.seqnum, bl)
        self.seqnum += 1
        self.mm.write(hdr)
        self.mm.write(blob)
        remaining = self.mm.size() - self.mm.tell()
        if remaining <= 0:
            self.close()
        return remaining

