package diagnostics

// TraceFunc trace function handler.
type traceFunc func(interface{}, error)

// use to store the trace hooks.
var listeners map[string]traceFunc

func init() {
	listeners = make(map[string]traceFunc)
	innerRegister()
}

// IsEnabled check the event whether has been registered.
func IsEnabled(name string) bool {
	_, ok := listeners[name]
	return ok
}

// Write trigger the event.
func Write(name string, obj interface{}, err error) {
	t, ok := listeners[name]
	if ok {
		// recover from panic when write error.
		defer func() {
			if r := recover(); r != nil {
				// suppress
			}
		}()

		t(obj, err)
	}
}

// WriteIf trigger the event if the name has been registered.
func WriteIf(name string, obj interface{}, err error) {
	t, ok := listeners[name]
	if ok {
		// recover from panic when write error.
		defer func() {
			if r := recover(); r != nil {
				// suppress
			}
		}()

		t(obj, err)
	}
}

// RegisterFunc register a func.
func RegisterFunc(name string, f func(interface{}, error)) {
	if f == nil {
		panic("[diagnostics] the register function is nil")
	}
	listeners[name] = traceFunc(f)
}
