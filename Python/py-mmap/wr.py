#
from common import MMJournal

## ############################################################

FILE_SIZE = 1024 * 1024 # (1 Meg)

TEST_FILE = '/tmp/test-py-mmap.out'

if __name__=='__main__':

    print("File size: {}".format(FILE_SIZE))
    num_msgs = FILE_SIZE / (14 + 15)
    print("Num msgs:  {}".format(num_msgs))

    j = MMJournal(TEST_FILE, FILE_SIZE)
    j.append('stub-data')
    j.append('stub-2')
    j.append('stub-record-3')
    try:
        for ii in range(num_msgs):
            j.append("stub-data-{}".format(ii))
    except:
        print("exception!")
        print("ii = {}".format(ii))
    finally:
        rem = j.mm.size() - j.mm.tell()
        print("remaining: {}".format(rem))
        j.close()
    print('done.')
