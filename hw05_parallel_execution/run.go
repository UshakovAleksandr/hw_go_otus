package hw05parallelexecution

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// wg
	wg := sync.WaitGroup{}
	// в него пишем таски
	chIn := make(chan func() error)
	// из него читаем ошибки для коллекции
	chErr := make(chan error)
	// тех канал. Из него вычитывают nil все горутины после max err
	chDone := make(chan struct{})

	// цикл запуска воркеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// воркеры по очереди вытаскивают таски из chIn
			for task := range chIn {
				// проверяем таску на err
				if err := task(); err != nil {
					// если err, то пишем в chErr
					select {
					case chErr <- err:
					// когда max err, работающие воркеры вычитывают
					// из канала nil и останавливаются.
					// В канала nil, тк нет писателей
					case <-chDone:
						return
					}
				}
			}
		}()
	}
	// пишем таски в канал
	go func() {
		for _, v := range tasks {
			chIn <- v
		}
		// закрываем для range
		close(chIn)
	}()

	// запускаем в отдельной горутине для закрытия канала chErr
	// после остановки воркеров
	go func() {
		wg.Wait()
		close(chErr)
	}()

	maxErrorsCount := m
	var counter int

	wg.Wait()
	// вычитываем из chErr без дедлока, тк выше wg.Wait() запущена в горутине
	for i := range chErr {
		fmt.Println(i)
		counter++
		if counter == maxErrorsCount {
			// закрываем тех канал
			close(chDone)
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}

// делал на этом примере, randErr() использовал для рандомной ошибки
// мысль была такая

// type Task func() error
//
// func randErr() error {
//	rand.Seed(time.Now().UnixNano())
//	a := rand.Intn(2-0) + 0
//	if a == 0 {
//		return fmt.Errorf("%v", "ошибка")
//	}
//	return nil
//}
//
// func Run() {
//	var tasks []Task
//	for i := 0; i < 10; i++ {
//		tasks = append(tasks, func() error {
//			return randErr()
//		})
//	}
//
//	wg := sync.WaitGroup{}
//	chIn := make(chan func() error)
//	chErr := make(chan error)
//	chDone := make(chan struct{})
//
//	for i := 0; i < 4; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for task := range chIn {
//				if err := task(); err != nil {
//					select {
//					case chErr <- err:
//					case <-chDone:
//						return
//					}
//				}
//			}
//		}()
//	}
//
//	go func() {
//		for _, v := range tasks {
//			chIn <- v
//		}
//		close(chIn)
//	}()
//
//	// wg дожедается отраблтки всех воркеров
//	go func() {
//		wg.Wait()
//		close(chErr)
//	}()
//
//	a := 4
//	var counter int
//
//	for i := range chErr {
//		fmt.Println(i)
//		counter++
//		if counter == a {
//			close(chDone)
//			break
//		}
//	}
//	// проверял второй wg.Wait() точно завершились ли все горутины, если не висит, то ок
//	//wg.Wait()
//}
