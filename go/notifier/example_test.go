package notifier_test

import (
	"context"
	"fmt"
	"time"

	"github.com/egustafson/sandbox/go/notifier"
)

func noticeListener(ch <-chan notifier.Notice) {
	fmt.Println("listener: started")
	for notice := range ch {
		msg, ok := notice.(string)
		if !ok {
			fmt.Println("listener: received non-string notice")
		} else {
			fmt.Printf("listener: %s\n", msg)
		}
	}
	fmt.Println("listener: exiting")
}

func noticeCallback(notice notifier.Notice) {
	msg, ok := notice.(string)
	if !ok {
		fmt.Println("cb: received non-string notice")
	} else {
		fmt.Printf("cb: %s\n", msg)
	}
}

func Example() {
	fmt.Println("")
	ctx, cancel := context.WithCancel(context.Background())
	// no defer cancel() - we explicitly cancel() in this func

	n := notifier.NewNotifier()

	n.RegCallback(ctx, noticeCallback)
	n.Notify("[ only callback registered ]")

	go noticeListener(n.ListenerChan(ctx))
	n.Notify("[ both cb and chan registered ]")

	go noticeListener(n.ListenerChan(context.Background())) // another listener, no ctx cancel
	n.Notify("[ second notice ]")

	time.Sleep(time.Millisecond) // let things settle a bit
	cancel()
	n.Notify("[ post cancel() notice ]")

	n.Close() // show Close() can be invoked after all listeners/cb's are canceled.

	time.Sleep(time.Millisecond) // let things settle a bit
	fmt.Println("Example() ending")
	// Unordered output:
	// cb: [ only callback registered ]
	// cb: [ both cb and chan registered ]
	// listener: started
	// listener: [ both cb and chan registered ]
	// listener: started
	// listener: [ second notice ]
	// listener: [ second notice ]
	// cb: [ second notice ]
	// listener: exiting
	// listener: [ post cancel() notice ]
	// listener: exiting
	// Example() ending
}

// ExampleBlocking demonstrates the case where the listener chan
// blocks long enough to be "problematic" and the Notify() routine
// spins the notice out into a separate goroutine.
func Example_blocking() {
	n := notifier.NewNotifier()
	ctx := context.Background()

	ch := n.ListenerChan(ctx) // get the chan but don't consume from it yet.

	n.Notify("[ msg 1 ]")
	n.Notify("[ msg 2 ]") // <-- this message should spawn a goroutine
	n.Notify("[ msg 3 ]") // <-- another goroutine ==> possible out of order deivery

	go noticeListener(ch)        // now the messages drain, and print
	time.Sleep(time.Millisecond) // let things settle
	n.Close()
	time.Sleep(time.Millisecond) // let things settle, again
	fmt.Println("ExampleBlocking() ending")
	// Unordered output:
	// timeout: deferring notice
	// timeout: deferring notice
	// listener: started
	// listener: [ msg 1 ]
	// listener: [ msg 2 ]
	// listener: [ msg 3 ]
	// listener: exiting
	// ExampleBlocking() ending
}
