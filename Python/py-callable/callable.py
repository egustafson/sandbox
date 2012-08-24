

class MyCall:
    def __call__(self):
        print "Called callable method."
        return


def func():
    print "Called function func()."

cbList = []
cbList.append( MyCall() )
cbList.append( func )

for cb in cbList:
    cb()

