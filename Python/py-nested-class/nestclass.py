import threading

class Master:

    class Monitor(threading.Thread):
        def run(self):
            print "Monitor thread has run."
            return

    def __init__(self):
        self.monitor = Master.Monitor()
        return

    def start(self):
        self.monitor.start()
        return

if __name__ == '__main__':
    m = Master()
    m.start()
