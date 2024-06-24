package marshal

import "encoding/json"

type I interface {
	Method()
}

type WI struct {
	I I
}

type WOI struct {
	I int
}

func M() {
	var i I
	json.Marshal(i) // want `interface value marshal.I is exported as json`

	wi := WI{}
	json.Marshal(wi) // want `interface field I is exported as json attribute`

	woi := WOI{}
	json.Marshal(woi)

	json.Marshal(struct {
		X int
	}{1})

	json.Marshal(struct { // want `interface field X is exported as json attribute`
		X I
	}{nil})

	json.NewEncoder(nil).Encode(i)
	json.NewEncoder(nil).Encode(wi)
	json.NewEncoder(nil).Encode(woi)
}
