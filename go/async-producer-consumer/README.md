Buffered, Async Producer - Consumer
===================================

This example shows a disconnected producer and consumer:

```
     ----------       --------       ----------
    | Consumer | <-- | Buffer | <-- | Producer |
     ----------       --------       ----------
```

Goal:  Never block the Producer or Consumer IF there is a message
available.

The Producer and Consumer can work at different rates.  The size of
the buffer is reported and it would be possible to mitigate if the
size exceeded some threshold.

The Buffer can be Close()ed allowing the Producer to signal that
it is done producing.

The Consumer receives both the next value and a 'more' flag indicating
that there is potentially more data coming.  Only when the producer
Close()'s the buffer does it return more=false to the consumer.

