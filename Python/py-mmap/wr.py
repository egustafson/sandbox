from mmap import mmap

from struct import pack


## Notes:  Create a publishable (and subscribable) object that wraps an mmap.
##  * publish -> (binary)Length + string
##  * ? fixed length file & when pub crosses boundary, create next file?
##  ?? how to open a new mm file and grow it with out stream writing (ie zero size problem)
##  Create a reader example too (same example file, use click?


FILE_SIZE = 1024 * 1024 * 1024  # (1 Gig)

class JournalConsumer(object):

    def __init__(self, journal):
        self._j = journal
        self._pos = 0

    def seek(self, index):
        None

    def next(self):
        None  ## (return next blob)


class MMJournal(object):

    def __init__(self, filename):
        self.filename = filename
        self.seqnum = 0
        self.f = open(filename, "w+b")
        self.f.write('\0')
        self.f.flush()
        self.mm = mmap(self.f.fileno(), 0)
        self.mm.resize(FILE_SIZE)
        self.mm.seek(0)


    def close(self):
        self.mm.flush()
        self.mm.close()
        self.f.close()


    def append(self, blob):
        bl = len(blob)
        end = self.mm.tell() + 12 + bl
        if end > self.mm.size():
            raise "exceeded file size"
        hdr = pack('cchll', 'j', 'l', 0, self.seqnum, bl)
        self.seqnum += 1
        self.mm.write(hdr)
        self.mm.write(blob)

    def getConsumer(self):
        return JournalConsumer(self)


if __name__=='__main__':

    j = MMJournal('data.out')
    j.append('stub-data')
    j.append('stub-2')
    j.append('stub-record-3')
    for ii in range(10000000):
        j.append("stub-data-{}".format(ii))
    rem = j.mm.size() - j.mm.tell()
    print("remaining: {}".format(rem))
    j.close()
    print('done.')


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
