#
from common import JournalConsumer

import common

## ############################################################

TEST_FILE = '/tmp/test-py-mmap.out'

if __name__=='__main__':

    print("header size: {}".format(common.HDR_SIZE))

    c = JournalConsumer(TEST_FILE)
    for ii in range(10):
        ( seq, blob ) = c.next()
        print("{}: {}".format(seq, blob))

    print('done.')
