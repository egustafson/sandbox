#
import threading
import time

## ############################################################

# Meta-resource -- in reality there is some resource being
# updated and _then_ the sequence (seq) counter is
# semi-atomically updated.
#
# Assumption:  a SINGLE updating thread.
#

class Resource(object):

    def __init__(self):
        self._cond = threading.Condition()
        self._blocks = 0
        self._seq = 0

    #
    # Only block to update the seq reference
    #
    def update(self):
        with self._cond:
            self._seq += 1
            self._cond.notify_all()
        print("update({})".format(self._seq))

    #
    # Only block if the seq has not been added (i.e. seq < idx)
    #
    def get(self, idx):
        if self._seq < idx:
            #print('block')
            with self._cond:
                self._blocks += 1
                while self._seq < idx:
                    self._cond.wait()
        return idx

    def get_blocks(self):
        return self._blocks


class Consumer(threading.Thread):

    def __init__(self, resource):
        threading.Thread.__init__(self)
        self._resource = resource

    def run(self):
        idx = 0
        t_ident = self.ident
        while idx < 10:
            idx += 1
            self._resource.get(idx)
            print("[{}] - {}".format(t_ident, idx))
        print("[{}] - thread-complete".format(t_ident))


## Main ###################################

if __name__=='__main__':

    num_consumers = 10
    count_to = 10

    consumers = []
    r = Resource()
    for x in range(0,num_consumers):
        c = Consumer(r)
        c.start()
        consumers.append(c)

    for x in range(0,count_to):
        r.update()
        #time.sleep(0.001)

    print('joining.')
    for c in consumers:
        c.join()
    print("Blocked {} times.".format(r.get_blocks()))
    print('done.')


## Local Variables:
## mode: python
## End:
