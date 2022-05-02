package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	return stages[3](stages[2](stages[1](stages[0](in))))

	// первый тест проходит
	// вопрос по второму, не могу понять как сделать, делаю как написал ниже и проблема в то, что для проверки done нужно использовать
	// select и соответственно горутину, но горутина не имеент return значения и поэтому пишу в канал - ch <- stages[3](stages[2](stages[1](stages[0](in)))),
	// и получается, что я пишу в канал объект другого канала, и закономерно падаю на result = append(result, s.(string)) в тесте, потому что там не строка, а указаетель на канал
	// как сделать? никак не могу додуматься

	// и более того, не могу понять как сделать вызов стейдже по их количеству, а не хардкодить по индексам

	//result := func() Out {
	//	ch := make(Bi)
	//	go func() {
	//		defer close(ch)
	//		select {
	//		case <- done:
	//			return
	//		case ch <- stages[3](stages[2](stages[1](stages[0](in)))):
	//		}
	//	}()
	//	return ch
	//}
	//
	//return result()
}
