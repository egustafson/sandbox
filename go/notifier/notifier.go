package notifier

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Notice interface{}

type NotifyCallbackFn func(notice Notice)

type Notifier interface {
	ListenerChan(ctx context.Context) <-chan Notice
	RegCallback(ctx context.Context, callback NotifyCallbackFn)

	Notify(notice Notice)
	Close()
}

type notifyReceiver struct {
	ctx    context.Context
	cancel context.CancelFunc
	ch     chan<- Notice
}

type notifier struct {
	lock    sync.Mutex
	targets map[*notifyReceiver]struct{}
}

func NewNotifier() Notifier {
	return &notifier{
		targets: make(map[*notifyReceiver]struct{}),
	}
}

func (n *notifier) Close() {
	n.lock.Lock()
	defer n.lock.Unlock()

	for tgt := range n.targets {
		tgt.cancel()
		close(tgt.ch)
		delete(n.targets, tgt)
	}
}

func (n *notifier) Notify(notice Notice) {
	const timeout = time.Millisecond
	n.lock.Lock()
	defer n.lock.Unlock()

	for tgt := range n.targets {
		if tgt.ctx.Err() != nil {
			close(tgt.ch)
			delete(n.targets, tgt)
			continue
		}
		select {
		case <-tgt.ctx.Done():
			close(tgt.ch)
			delete(n.targets, tgt)
		case tgt.ch <- notice:
			// default case, just deliver
		case <-time.After(timeout):
			go func() { // spin delivery out into a goroutne, avoid blocking
				select {
				case <-tgt.ctx.Done():
					return
				case tgt.ch <- notice:
				}
			}()
			fmt.Println("timeout: deferring notice")
		}
	}
}

func (n *notifier) ListenerChan(ctx context.Context) <-chan Notice {
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan Notice, 1)
	tgt := &notifyReceiver{
		ctx:    ctx,
		cancel: cancel,
		ch:     ch,
	}
	n.lock.Lock()
	defer n.lock.Unlock()

	n.targets[tgt] = struct{}{}
	return ch
}

func (n *notifier) RegCallback(ctx context.Context, callback NotifyCallbackFn) {
	listenChan := n.ListenerChan(ctx)

	// start a goroutine that calls callback when a notice arrives on
	// the channel.  This func exits when ctx is canceled because
	// Notify() closes channels with canceled context's.
	go func() {
		for notice := range listenChan {
			callback(notice)
		}
	}()
}
