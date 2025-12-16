package abstractions

type Counter struct {
	value uint32
}

func NewCounter() *Counter {
	return &Counter{value: 0}
}

func (c *Counter) Increment() {
	c.value++
}

func (c *Counter) GetValue() uint32 {
	return c.value
}

func (c *Counter) Reset() {
	c.value = 0
}
