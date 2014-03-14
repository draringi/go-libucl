package libucl

// #include <ucl.h>
import "C"

// Object represents a single object within a configuration.
type Object struct {
	object *C.ucl_object_t
}

type ObjectType int

const (
	ObjectTypeObject ObjectType = iota
	ObjectTypeArray
	ObjectTypeInt
	ObjectTypeFloat
	ObjectTypeString
	ObjectTypeBoolean
	ObjectTypeTime
	ObjectTypeUserData
	ObjectTypeNull
)

// Free the memory associated with the object. This must be called when
// you're done using it.
func (o *Object) Close() {
	C.ucl_object_unref(o.object)
}

func (o *Object) Get(key string) *Object {
	obj := C.ucl_object_find_keyl(o.object, C.CString(key), C.size_t(len(key)))
	if obj == nil {
		return nil
	}

	return &Object{object: obj}
}

// Returns the key of this value/object as a string, or the empty
// string if the object doesn't have a key.
func (o *Object) Key() string {
	return C.GoString(C.ucl_object_key(o.object))
}

// Len returns the length of the object, or how many elements are part
// of this object.
//
// For objects, this is the number of key/value pairs.
// For arrays, this is the number of elements.
func (o *Object) Len() uint {
	return uint(o.object.len)
}

// Returns the type that this object represents.
func (o *Object) Type() ObjectType {
	return ObjectType(o.object._type)
}

//------------------------------------------------------------------------
// Conversion Functions
//------------------------------------------------------------------------

func (o *Object) ToString() string {
	return C.GoString(C.ucl_object_tostring(o.object))
}