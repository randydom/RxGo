package rxgo

type (
	// NextFunc handles a next item in a stream.
	NextFunc func(interface{})

	// ErrFunc handles an error in a stream.
	ErrFunc func(error)

	// DoneFunc handles the end of a stream.
	DoneFunc func()
)

// EventHandler type is implemented by all handlers and Observer.
type EventHandler interface {
	Handle(interface{})
}

// Handle registers NextFunc to EventHandler.
func (handle NextFunc) Handle(item interface{}) {
	switch item := item.(type) {
	case error:
		return
	default:
		handle(item)
	}
}

// Handle registers ErrFunc to EventHandler.
func (handle ErrFunc) Handle(item interface{}) {
	switch item := item.(type) {
	case error:
		handle(item)
	default:
		return
	}
}

// Handle registers DoneFunc to EventHandler.
func (handle DoneFunc) Handle(item interface{}) {
	handle()
}
