package goroutineutils

type GoroutinePool struct {
	pool chan struct{}
}

func NewGoroutinePool(maxWorkers int) *GoroutinePool {
	pool := make(chan struct{}, maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		pool <- struct{}{}
	}

	return &GoroutinePool{pool: pool}
}

func (p *GoroutinePool) Execute(fn func()) {
	<-p.pool
	go func() {
		fn()
		p.pool <- struct{}{}
	}()
}
