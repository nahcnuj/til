package synchronous

import "sync"

type Counter struct {
	mu    sync.Mutex // can be embedded, but which confuses because Lock & Unlock get callable pubilcly
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
