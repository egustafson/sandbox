#
from common import MMJournal

## ############################################################

MEGABYTE = 1024 * 1024

FILE_SIZE = 100 * MEGABYTE

TEST_FILE = '/tmp/test-py-mmap.out'

if __name__=='__main__':

    print("Target size: {}".format(FILE_SIZE))

    j = MMJournal(TEST_FILE, FILE_SIZE)
    j.append('stub-data')
    j.append('stub-2')
    j.append('stub-record-3')

    remaining = j.available()
    ii = 0
    while remaining > 0:
        remaining = j.append("stub-data-{}".format(ii))
        ii += 1

    print('--')
    print("Final size:  {}".format(j.size()))
    print("Entries:     {}".format(j.entries()))
    print('done.')
