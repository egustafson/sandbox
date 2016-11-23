#
from common import JournalConsumer

import common

## ############################################################

TEST_FILE = '/tmp/test-py-mmap.out'

if __name__=='__main__':

    print("header size: {}".format(common.HDR_SIZE))

    c = JournalConsumer(TEST_FILE)
    while not c.eof():
        (seq, blob) = c.next()

    print('--')
    print("Read {} bytes.".format(c.pos()))
    print("Read {} objects.".format(c.seq()))
    print('done.')
