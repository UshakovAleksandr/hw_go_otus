package hw06pipelineexecution

import (
	"sort"
	"sync"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	ch := make(Bi)
	tempCh := make(Bi)
	out := make(Bi)
	wg := sync.WaitGroup{}

loop:
	for {
		if done != nil {
			time.Sleep(time.Millisecond * 210)
		}
		select {
		case <-done:
			//fmt.Println("done")
			break loop
		default:
			i, ok := <-in
			if !ok {
				break loop
			}
			wg.Add(1)
			go func() {
				for i := range stages[3](stages[2](stages[1](stages[0](ch)))) {
					tempCh <- i
					wg.Done()
					//fmt.Println("worker")
				}
			}()
			ch <- i
		}
	}

	go func() {
		wg.Wait()
		close(tempCh)
	}()

	var resultSl []string

	for item := range tempCh {
		resultSl = append(resultSl, item.(string))
	}

	sort.Strings(resultSl)

	go func() {
		defer close(out)
		for _, item := range resultSl {
			out <- item
		}
	}()

	return out
}
