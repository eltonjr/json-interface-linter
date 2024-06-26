package marshalcustom

import (
	"context"
)

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
	wi := WI{}
	Encode(context.Background(), wi) // want `interface field I is exported as json attribute`

	woi := WOI{}
	Encode(context.Background(), woi)
}

func Encode(ctx context.Context, v any) {
	// let say it will encode to json
}
