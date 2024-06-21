# JSON interface linter

Uses go ast to check if serialized structs contain an interface.

Walks through every struct to check if it has any json tag.  
If it does, recursively checks every field to make sure none is an interface.

