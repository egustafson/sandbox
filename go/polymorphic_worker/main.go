package main

import "fmt"

type PoolWorker interface {
	DoWork()
}

type multPoolWorker struct {
	x, y     int
	resultCh chan int
	done     bool
}

type catPoolWorker struct {
	a, b     string
	resultCh chan string
	done     bool
}

var _ PoolWorker = (*multPoolWorker)(nil)
var _ PoolWorker = (*catPoolWorker)(nil)

func NewMultPoolWorker(x, y int, results chan int, done bool) PoolWorker {
	return &multPoolWorker{
		x:        x,
		y:        y,
		resultCh: results,
		done:     done,
	}
}

func NewCatPoolWorker(a, b string, results chan string, done bool) PoolWorker {
	return &catPoolWorker{
		a:        a,
		b:        b,
		resultCh: results,
		done:     done,
	}
}

func (w *multPoolWorker) DoWork() {
	r := w.x * w.y
	w.resultCh <- r
	if w.done {
		close(w.resultCh)
	}
}

func (w *catPoolWorker) DoWork() {
	r := fmt.Sprintf("%s%s", w.a, w.b)
	w.resultCh <- r
	if w.done {
		close(w.resultCh)
	}
}

func WorkerPool(ch chan PoolWorker) {
	fmt.Println("WorkerPool started.")
	for w := range ch {
		w.DoWork()
	}
	fmt.Println("WorkerPool is done.")
}

func main() {
	workerCh := make(chan PoolWorker)
	multResults := make(chan int)
	catResults := make(chan string)

	go WorkerPool(workerCh)

	go func() {
		for ir := range multResults {
			fmt.Printf("multResult: %d\n", ir)
		}
	}()

	go func() {
		for sr := range catResults {
			fmt.Printf("catResult: %s\n", sr)
		}
	}()

	for i := 1; i < 10; i++ {
		workerCh <- NewMultPoolWorker(i, i*2, multResults, false)
		workerCh <- NewCatPoolWorker(fmt.Sprintf("%d", i), "-tag", catResults, false)
	}
	workerCh <- NewMultPoolWorker(1, 1, multResults, true)
	workerCh <- NewCatPoolWorker("", "", catResults, true)

	close(workerCh)
	fmt.Println("DONE.")
}
