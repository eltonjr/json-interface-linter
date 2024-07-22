package structtag

import "io"

type I interface {
	ExportedMethod()
}

type AliasInterface I
type AliasInterface2 AliasInterface

type S struct {
}

type ExportedInterface struct {
	InterfaceField       I               `json:"i"`               // want `interface jsontag.I is exported as json`
	InterfaceStdField    io.Reader       `json:"reader"`          // want `interface io.Reader is exported as json`
	AliasInterfaceField  AliasInterface  `json:"alias"`           // want `interface jsontag.AliasInterface is exported as json`
	AliasInterface2Field AliasInterface2 `json:"alias2"`          // want `interface jsontag.AliasInterface2 is exported as json`
	Any                  any             `json:"any"`             // want `interface any is exported as json`
	EmptyInterface       interface{}     `json:"empty_interface"` // want `interface interface{} is exported as json`

	IntField    int `json:"exported_field"`
	StructField S   `json:"exported_struct"`

	Skippedfield    I `json:"-"`
	hiddenInterface I `json:"i"`
}

// If some field from struct is being exported, it means the whole struct will be exported
type HiddenInterface struct {
	InterfaceField          I   // want `interface jsontag.I is exported as json`
	TaggedNonInterfaceField int `json:"busted"`
}

type Constraint[T comparable] struct {
	T T `json:"t"` // want `interface T is exported as json`
}

type Recursive struct {
	I I `json:"i"` // want `interface jsontag.I is exported as json`
	R *Recursive
}

type UnexportedInterface struct {
	I I
}

type OtherTagsThatNotJSON struct {
	I I `something`
}
