package event

import "fmt"

// EventHandler struct declare
type EventHandler struct {
	listeners map[string]chan *EventArgs
}

// EventArgs struct declare
type EventArgs struct {
	Param map[string]interface{}
}

// NewEventHandler handler
func NewEventHandler() *EventHandler {
	return &EventHandler{make(map[string]chan *EventArgs)}
}

// NewEventArgs args
func NewEventArgs() *EventArgs {
	return &EventArgs{make(map[string]interface{})}
}

// AddListener callback
func (handler *EventHandler) AddListener(callbackName string, callback func(*EventArgs)) {
	if _, ok := handler.listeners[callbackName]; !ok {
		ch := make(chan *EventArgs)

		handler.listeners[callbackName] = ch
		go func() {
			for {
				event := <-ch
				if event == nil {
					break
				}
				(callback)(event)
			}
		}()
	}
}

// RemoveListener callback
func (handler *EventHandler) RemoveListener(callbackName string) (listenersLength int) {
	if _, ok := handler.listeners[callbackName]; ok {
		handler.listeners[callbackName] <- nil
		delete(handler.listeners, callbackName)
	}
	return len(handler.listeners)
}

// OnEvent event
func (handler *EventHandler) OnEvent(event *EventArgs) {
	for _, v := range handler.listeners {
		v <- event
	}
}

// Release handler
func (handler *EventHandler) Release() {
	for _, v := range handler.listeners {
		v <- nil
	}
}

func Test_Event() {
	handler := NewEventHandler()
	defer handler.Release()

	aaa := NewMouse("aaa")
	ccc := NewMouse("ccc")

	handler.AddListener("aaa", aaa.Shout)
	handler.AddListener("ccc", ccc.Shout)

	Alert()

	args := NewEventArgs()
	args.Param["info"] = "Run..."

	handler.OnEvent(args)

	fmt.Println("aaa is always eaten by big cat...")

	Alert()

	handler.RemoveListener("aaa")
	handler.OnEvent(args)
}

func Alert() {
	fmt.Println("Cat coming...")
}

func NewMouse(name string) *Mouse {
	fmt.Println("小老鼠", name, "出来了。")
	return &Mouse{name}
}

type Mouse struct {
	Name string
}

func (mouse *Mouse) Shout(args *EventArgs) {
	if v, ok := args.Param["info"]; ok {
		fmt.Println(mouse.Name, v.(string))
	}
}
