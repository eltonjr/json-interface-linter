package marshal

import "encoding/json"

type I interface {
	Method()
}

type I2 interface {
	Method()
}

type WI struct {
	I I
}

type W2I struct {
	I  I
	I2 I2
}

type WOI struct {
	I int
}

func M() {
	var i I
	json.Marshal(i) // want `interface marshal.I is exported as json`

	wi := WI{}
	json.Marshal(wi) // want `interface marshal.I is exported as json`

	w2i := W2I{}
	json.Marshal(w2i) // want `interface marshal.I is exported as json` `interface marshal.I2 is exported as json`

	woi := WOI{}
	json.Marshal(woi)

	json.Marshal(struct {
		X int
	}{1})

	json.Marshal(struct { // want `interface marshal.I is exported as json`
		X I
	}{nil})

	json.NewEncoder(nil).Encode(i)  // want `interface marshal.I is exported as json`
	json.NewEncoder(nil).Encode(wi) // want `interface marshal.I is exported as json`
	json.NewEncoder(nil).Encode(woi)
}
