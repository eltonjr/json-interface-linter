package structtag

type I interface {
	ExportedMethod()
}

type S struct {
}

type ExportedInterface struct {
	Ifield I `json:"i"` // want `interface field Ifield is exported as json attribute`
}

type ExportedField struct {
	ExportedField int `json:"exported_field"`
}

type ExportedStruc struct {
	ExportedStruc S `json:"exported_struct"`
}

type UnexportedInterface struct {
	I I
}
