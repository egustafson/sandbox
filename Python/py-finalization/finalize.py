#!/usr/bin/env python


class Finalizer:
    def __init__(self, id):
        self.id = id
        print "Birth: Finalizer(%s)" % self.id
        return

    def __del__(self):
        print "Death: Finalizer(%s)" % self.id
        return


def create(id):
    f = Finalizer(id)

if __name__ == "__main__":
    print "Begin..."
    f = Finalizer(1)    # Destroyed after 'done'
    create(2)           # destroyed before 3
    Finalizer(3)        # destroyed before 4
    create(4)           # destroyed before 'done'
    print "done."
    
