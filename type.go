package gopkg

import (
	"reflect"
	"unsafe"
)

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

// Ptr is the old name for the Pointer kind.
const Ptr = Pointer

var kindNames = []string{
	Invalid:       "invalid",
	Bool:          "bool",
	Int:           "int",
	Int8:          "int8",
	Int16:         "int16",
	Int32:         "int32",
	Int64:         "int64",
	Uint:          "uint",
	Uint8:         "uint8",
	Uint16:        "uint16",
	Uint32:        "uint32",
	Uint64:        "uint64",
	Uintptr:       "uintptr",
	Float32:       "float32",
	Float64:       "float64",
	Complex64:     "complex64",
	Complex128:    "complex128",
	Array:         "array",
	Chan:          "chan",
	Func:          "func",
	Interface:     "interface",
	Map:           "map",
	Pointer:       "ptr",
	Slice:         "slice",
	String:        "string",
	Struct:        "struct",
	UnsafePointer: "unsafe.Pointer",
}

type Typer interface {
	Type() string
}

// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i any) string {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	if int(eface.typ.kind) < len(kindNames) {
		return kindNames[eface.typ.kind]
	}
	return reflect.TypeOf(i).String()
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ *rtype
	_   unsafe.Pointer
}

// rtype is the common implementation of most values.
// It is embedded in other struct types.
//
// rtype must be kept in sync with ../runtime/type.go:/^type._type.
type rtype struct {
	_    uintptr
	_    uintptr // number of bytes in the type that can contain pointers
	_    uint32  // hash of type; avoids computation in hash tables
	_    uint8   // extra type information flags
	_    uint8   // alignment of variable with this type
	_    uint8   // alignment of struct field with this type
	kind uint8   // enumeration for C
	_    func(unsafe.Pointer, unsafe.Pointer) bool
	_    *byte // garbage collection data
	_    int32 // string form
	_    int32 // type for pointer to this type, may be zero
}

func GetType[T any](t T) string {
	return TypeOf(any(t))
}
