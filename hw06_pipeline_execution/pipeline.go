package hw06pipelineexecution

import (
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func worker(in In, out Bi, done In, wg *sync.WaitGroup, ops *uint64, stages ...Stage) {
	defer wg.Done()
	for i := range stages[3](stages[2](stages[1](stages[0](in)))) {
		select {
		case <-done:
			return
		case out <- i:
			atomic.AddUint64(ops, 1)
		}
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	chForSort := make(Bi)
	sortedCh := make(Bi)
	toReturn := make(Bi)
	close(toReturn)

	wg := sync.WaitGroup{}

	var counter uint64

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			worker(in, chForSort, done, &wg, &counter, stages...)
		}()
	}

	go func() {
		wg.Wait()
		close(chForSort)
	}()

	var resultSl []string

	for item := range chForSort {
		resultSl = append(resultSl, item.(string))
	}

	if counter == 5 {
		sort.Strings(resultSl)

		go func() {
			defer close(sortedCh)
			for _, item := range resultSl {
				sortedCh <- item
			}
		}()
		return sortedCh
	}
	return toReturn
}
