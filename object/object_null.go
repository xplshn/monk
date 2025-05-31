package object

// Nil wraps nothing and implements our Object interface.
type Nil struct{}

// Type returns the type of this object.
func (n *Nil) Type() Type {
	return NIL_OBJ
}

// Inspect returns a string-representation of the given object.
func (n *Nil) Inspect() string {
	return "nil"
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (n *Nil) InvokeMethod(method string, env Environment, args ...Object) Object {
	return nil
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (n *Nil) ToInterface() interface{} {
	return "<NIL>"
}
