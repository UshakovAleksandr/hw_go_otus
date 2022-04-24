package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n, m int) error {
	// define chans
	chIn := make(chan Task)
	chErr := make(chan error)
	chTooManyErrs := make(chan struct{})
	chAllDone := make(chan struct{})

	// run workers
	wg := sync.WaitGroup{}
	defer wg.Wait()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range chIn {
				if err := task(); err != nil {
					select {
					case chErr <- err:
					case <-chTooManyErrs:
						return
					}
				}
			}
		}()
	}

	// close chAllDone when all workers exited
	go func() {
		wg.Wait()
		close(chAllDone)
	}()

	// feed input channel with tasks
	go func() {
		defer close(chIn)
		for _, task := range tasks {
			select {
			case chIn <- task:
			case <-chTooManyErrs:
				// stop when too many errors
				return
			}
		}
	}()

	// count error
	errCounter := 0
	for {
		select {
		case <-chErr:
			errCounter++
			if errCounter >= m {
				close(chTooManyErrs)
				fmt.Printf("too many errs: %d\n", errCounter)
				return ErrErrorsLimitExceeded
			}
		case <-chAllDone:
			fmt.Println("all done with success")
			return nil
		}
	}
}
