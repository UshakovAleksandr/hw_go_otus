package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	ch := make(Bi)

	select {
	case <-done:
		close(ch)
		return ch
	default:
		return stages[3](stages[2](stages[1](stages[0](in))))
	}

	// изменил немного, ошибку предыдушего коммента понял
	// но второй тест по прежнему не могу пройти. Причины тоже понятны, неправильный select.
	// Если я правильно понимаю, то принцип должен быть такой:
	// перед select нужжен for{}, в котором постоянно проверяется передаваемое значение обоих каналов
	// записываем полученные из стейджей значения, сколько успеет выполнится и если done стал закрыт, то нужно прибить эти
	// значения и отправить закрытый канал
	//
	//for {
	//	select {
	//	case <-done:
	//		close(ch)
	//		return ch
	//	default:
	//		someCh <- stages[3](stages[2](stages[1](stages[0](in))))
	//      ...
	//	}
	//}
	// return ...
}
