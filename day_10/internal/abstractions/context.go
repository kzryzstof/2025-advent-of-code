package abstractions

type Context struct {
	isCompleted  bool
	buttonsCount int
}

func NewContext(
	buttonsCount int,
) *Context {
	return &Context{
		isCompleted:  false,
		buttonsCount: buttonsCount,
	}
}

func (c *Context) IsCompleted() bool {
	return c.isCompleted
}
