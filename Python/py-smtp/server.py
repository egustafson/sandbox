import asyncore
from smtpd import SMTPServer


class Server(SMTPServer):

    def process_message(self, peer, mailfrom, rcptos, data):
        print('-- new message --')
        print("peer:  {!r}".format(peer))
        print("from:  {!r}".format(mailfrom))
        print("to:    {!r}".format(rcptos))
        print("data:  {} bytes".format(len(data)))
        print("\n{}".format(data))
        print(".")


if __name__=='__main__':

    server = Server(('127.0.0.1', 1025), None)
    asyncore.loop()  # forever

