package work

import "sync"

//work必须满足接口类型才能使用工作池

type Worker interface {
	Task()
}

//pool 提供一个goroutine池，这个池可以完成
//任何已提交的Worker任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//new创建一个新工作池
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	//开启线程池去完成任务
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			//从通道中取出一个任务
			//循环终止条件
			//无缓冲通道，通道被关闭，循环终止
			for w := range p.work {
				//执行任务
				w.Task()
			}
			//当通道里面没有任务的时候，计数器减一
			p.wg.Done()
		}()
	}
	return &p
}

//Run提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//shutdown等待所有goroutine停止工作
func (p *Pool) Shutdown() {
	//关闭通道
	close(p.work)
	p.wg.Wait()
}
