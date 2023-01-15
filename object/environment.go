package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = env
	return env

}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name] // inner scope
	if !ok && e.outer != nil {
		obj, ok := e.outer.Get(name) // outer scope
		return obj, ok
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
