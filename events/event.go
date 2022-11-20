package events

// Event 事件传输结构
type Event struct {
	Name string
	Data interface{}
}

// NewEvent 构建新的事件结构实例
func NewEvent(name string, data interface{}) *Event {
	return &Event{Name: name, Data: data}
}
