package hw06pipelineexecution

import (
	"sync"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

// вот на мой взгляд максимально близкий вариант, но не могу понять, почему не закрывается канал
// go func() {
//		wg.Wait()
//		close(out)
//	}()
// не отрабатывает
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	ch := make(Bi)
	out := make(Bi)
	wg := sync.WaitGroup{}

	for i := range in {
		select {
		case <- done:
			return out
		default:
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := range stages[3](stages[2](stages[1](stages[0](ch)))) {
					out <- i
				}
			}()
			ch <- i
		}
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

//func worker1(in In, out Bi, done In, wg *sync.WaitGroup, ops *uint64, stages ...Stage) {
//	defer wg.Done()
//	for i := range stages[3](stages[2](stages[1](stages[0](in)))) {
//		select {
//		case <-done:
//			return
//		case out <- i:
//			atomic.AddUint64(ops, 1)
//		}
//	}
//}
//
//func ExecutePipeline1(in In, done In, stages ...Stage) Out {
//	chForSort := make(Bi)
//	sortedCh := make(Bi)
//	toReturn := make(Bi)
//	close(toReturn)
//
//	wg := sync.WaitGroup{}
//
//	var counter uint64
//
//	for i := 0; i < runtime.NumCPU(); i++ {
//		wg.Add(1)
//		go func() {
//			worker1(in, chForSort, done, &wg, &counter, stages...)
//		}()
//	}
//
//	go func() {
//		wg.Wait()
//		close(chForSort)
//	}()
//
//	var resultSl []string
//
//	for item := range chForSort {
//		resultSl = append(resultSl, item.(string))
//	}
//
//	if counter == 5 {
//		sort.Strings(resultSl)
//
//		go func() {
//			defer close(sortedCh)
//			for _, item := range resultSl {
//				sortedCh <- item
//			}
//		}()
//		return sortedCh
//	}
//	return toReturn
//}
