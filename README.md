# JSON interface linter

Uses go ast to check if serialized structs contain an interface.

Walks through every struct to check if it has any json tag.  
If it does, recursively checks every field to make sure none is an interface.

### Running

First you need to install it locally (this linter is not published in golangci-ci)

```
go install cmd/json-interface/json-interface.go
```

Then you can run it at the root of your package

```
cd <path-to-my-package>
json-interface  ./...
```

### TODO
- [ ] Check nested (recursive) fields
- [ ] If 1 field has json tag, assume the whole struct is exposed and check every other
- [x] skip `json:"-"`
- [x] json.Marshal
- [x] encoder.Encode
- [x] chi / gin
- [ ] Check if map has any interface
- [ ] Verbose mode
