package object

import (
	"fmt"
	"sort"
	"strings"
)

type (
	Int8    struct{ Value int8 }
	Int16   struct{ Value int16 }
	Int32   struct{ Value int32 }
	Int64   struct{ Value int64 }
	Int     struct{ Value int }
	Uint8   struct{ Value uint8 }
	Uint16  struct{ Value uint16 }
	Uint32  struct{ Value uint32 }
	Uint64  struct{ Value uint64 }
	Uint    struct{ Value uint }
	Uintptr struct{ Value uintptr }
)

func (i *Int8) Inspect() string    { return fmt.Sprintf("%d", i.Value) }
func (i *Int16) Inspect() string   { return fmt.Sprintf("%d", i.Value) }
func (i *Int32) Inspect() string   { return fmt.Sprintf("%d", i.Value) }
func (i *Int64) Inspect() string   { return fmt.Sprintf("%d", i.Value) }
func (i *Int) Inspect() string     { return fmt.Sprintf("%d", i.Value) }
func (i *Uint8) Inspect() string   { return fmt.Sprintf("%d", i.Value) }
func (i *Uint16) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Uint32) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Uint64) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Uint) Inspect() string    { return fmt.Sprintf("%d", i.Value) }
func (i *Uintptr) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Int8) Type() Type    { return INT8_OBJ }
func (i *Int16) Type() Type   { return INT16_OBJ }
func (i *Int32) Type() Type   { return INT32_OBJ }
func (i *Int64) Type() Type   { return INT64_OBJ }
func (i *Int) Type() Type     { return INT_OBJ }
func (i *Uint8) Type() Type   { return UINT8_OBJ }
func (i *Uint16) Type() Type  { return UINT16_OBJ }
func (i *Uint32) Type() Type  { return UINT32_OBJ }
func (i *Uint64) Type() Type  { return UINT64_OBJ }
func (i *Uint) Type() Type    { return UINT_OBJ }
func (i *Uintptr) Type() Type { return UINTPTR_OBJ }

func (i *Int8) HashKey() HashKey    { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Int16) HashKey() HashKey   { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Int32) HashKey() HashKey   { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Int64) HashKey() HashKey   { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Int) HashKey() HashKey     { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Uint8) HashKey() HashKey   { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Uint16) HashKey() HashKey  { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Uint32) HashKey() HashKey  { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Uint64) HashKey() HashKey  { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Uint) HashKey() HashKey    { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }
func (i *Uintptr) HashKey() HashKey { return HashKey{Type: i.Type(), Value: uint64(i.Value)} }

func (i *Int8) ToInterface() interface{}    { return i.Value }
func (i *Int16) ToInterface() interface{}   { return i.Value }
func (i *Int32) ToInterface() interface{}   { return i.Value }
func (i *Int64) ToInterface() interface{}   { return i.Value }
func (i *Int) ToInterface() interface{}     { return i.Value }
func (i *Uint8) ToInterface() interface{}   { return i.Value }
func (i *Uint16) ToInterface() interface{}  { return i.Value }
func (i *Uint32) ToInterface() interface{}  { return i.Value }
func (i *Uint64) ToInterface() interface{}  { return i.Value }
func (i *Uint) ToInterface() interface{}    { return i.Value }
func (i *Uintptr) ToInterface() interface{} { return i.Value }

func (i *Int8) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeIntegerMethods(int64(i.Value), "int8", method, env)
}
func (i *Int16) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeIntegerMethods(int64(i.Value), "int16", method, env)
}
func (i *Int32) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeIntegerMethods(int64(i.Value), "int32", method, env)
}
func (i *Int64) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeIntegerMethods(i.Value, "int64", method, env)
}
func (i *Int) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeIntegerMethods(int64(i.Value), "int", method, env)
}

func (i *Uint8) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeUnsignedMethods(uint64(i.Value), "uint8", method, env)
}
func (i *Uint16) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeUnsignedMethods(uint64(i.Value), "uint16", method, env)
}
func (i *Uint32) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeUnsignedMethods(uint64(i.Value), "uint32", method, env)
}
func (i *Uint64) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeUnsignedMethods(i.Value, "uint64", method, env)
}
func (i *Uint) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeUnsignedMethods(uint64(i.Value), "uint", method, env)
}
func (i *Uintptr) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "to_f" {
		return &Float{Value: float64(i.Value)}
	}
	return invokeUnsignedMethods(uint64(i.Value), "uintptr", method, env)
}

func invokeIntegerMethods(value int64, prefix, method string, env Environment) Object {
	switch method {
	case "chr":
		return &String{Value: string(rune(value))}
	case "methods":
		static := []string{"chr", "methods"}
		dynamic := env.Names(prefix + ".")
		var names []string
		names = append(names, static...)
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	default:
		return nil
	}
}

func invokeUnsignedMethods(value uint64, prefix, method string, env Environment) Object {
	switch method {
	case "methods":
		static := []string{"methods"}
		dynamic := env.Names(prefix + ".")
		var names []string
		names = append(names, static...)
		for _, e := range dynamic {
			bits := strings.Split(e, ".")
			names = append(names, bits[1])
		}
		sort.Strings(names)
		result := make([]Object, len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	default:
		return nil
	}
}
