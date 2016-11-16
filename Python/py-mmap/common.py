from binascii import crc32
from mmap import mmap, ACCESS_READ
from struct import pack, unpack, calcsize


## Notes:  Create a publishable (and subscribable) object that wraps an mmap.
##  * publish -> (binary)Length + string
##  * ? fixed length file & when pub crosses boundary, create next file?
##  ?? how to open a new mm file and grow it with out stream writing (ie zero size problem)
##  Create a reader example too (same example file, use click?


DEFAULT_SIZE = 1024 * 1024  # (1 Meg)

HDR_FMT  = '=iHII'  # crc32, ver, seq, size
HDR_SIZE = calcsize(HDR_FMT)

class JournalConsumer(object):

    def __init__(self, filename):
        self._pos = 0
        self._seq = 0
        self.filename = filename
        self.f = open(filename, "r")
        self.mm = mmap(self.f.fileno(), 0,
                       access=ACCESS_READ)

    def seek(self, index):
        None

    def next(self):
        hdr = self.mm.read(HDR_SIZE)
        (crc, ver, seq, size) = unpack(HDR_FMT, hdr)
        self._seq = seq
        blob = self.mm.read(size)
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
        self.mm.seek(0)


    def close(self):
        self.mm.flush()
        self.mm.close()
        self.f.close()


    def append(self, blob):
        bl = len(blob)
        crc = crc32(blob)
        end = self.mm.tell() + HDR_SIZE + bl
        if end > self.mm.size():
            raise "exceeded file size"
        ver = 0
        hdr = pack(HDR_FMT, crc, ver, self.seqnum, bl)
        self.seqnum += 1
        self.mm.write(hdr)
        self.mm.write(blob)


# with open('data.out', 'w+b') as f:
#     f.write("-")
#     f.flush()
#     mm = mmap(f.fileno(), 0)
#     if mm.size() < 1024:
#         print("resizing.")
#         mm.resize(1024)
#     mm.seek(0)
#     for ii in range(10):
#         mm.write("data-")
#     mm.flush()
#     mm.close()
#
# print('done.')
