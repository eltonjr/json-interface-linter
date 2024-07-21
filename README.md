# JSON interface linter

Checks a go module to ensure no data structure containing an interface is being marshaled as json.

This can be used to validate the suitability for a drop-in replacement of a json encoder.
Some of these encoders do not allow interfaces being marshaled, for performance reasons.

## Running

1. Install it locally

```
go install github.com/eltonjr/json-interface-linter@latest
```

1. Run it at the root of the module

Suggestion: don't validate test files

```
cd <path-to-my-module>
json-interface -test=false ./...
```

## The linter

This module provides two linters that can be used individually or together

* jsontag
* marshal

### jsontag

Checks every struct of a module looking for tagged fields `json:"..."`.
If some struct has a json tag, it is considered exported, and every field of this struct will be checked not to be an interface.

Fields unexported (starting with lower-case) or marked as `json:"-"` will not be considered.

This linter can be used individually with the following argument:

```
json-interface -jsontag ./...
```

### marshal

Checks every value being passed to a `Marshal` function does not contain an interface within it.

By default, it checks function calls to
```
json.Marshal         (from encode/json)
json.MarshalIndent   (from encode/json)
json.Encode          (from encode/json.Encoder)
```

Custom marshalers can also be checked by being specified in a file and informed to the linter

```
json-interface -marshal.marshalers=file.txt ./...
```

A custom marshaler can also inform which of the arguments is the value being marshaled.
Arguments are optional, positional and zero based, with zero as default.

Example of a marshalers file:
```
github.com/gin-gonic/gin.JSON
myencoder.Encode[1]
```

**Note:** providing a marshalers list will override the default marshalers. To be considered, they must also be in the marshalers file

More examples can be found [here](marshal/testdata/valid.txt)

This linter can be used individually with the following argument:

```
json-interface -marshal ./...
```

## Excluding false-positives

If false-positives were found, they can be ignored by being specified in a file and informed to the linter

```
json-interface -exclude=exclude.txt
```

Example of an exclusion file:
```
errors
any
mypkg.MyInterface
```

## Supported flags

| Flag                | Description                                                  | Default                                                                |
|---------------------|--------------------------------------------------------------|------------------------------------------------------------------------|
| -test               | Specifies whether to validate test files.                    | true                                                                   |
| -jsontag            | Runs the jsontag linter.                                     | true                                                                   |
| -marshal            | Runs the marshal linter.                                     | true                                                                   |
| -marshal.marshalers | Specifies a file containing custom marshalers to be checked. | encoding/json.Marshal encoding/json.MarshalIndent encoding/json.Encode |
| -exclude            | Specifies a file containing false-positives to be ignored.   | <empty>                                                                |
| -verbose            | Logs every decision made by the linter. Useful for debug.    | false                                                                  |