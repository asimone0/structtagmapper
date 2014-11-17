structtagmapper
===============

A utility for parsing a reflect.StructTag into map[string]string
Based on the source code for reflect.StructTag.Get(key string)

example usage:

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/tonestuff/structtagmapper"
)

type E struct {
	ID int `xml:"id" json:"id" comment:"test comment"`
}

func main() {
	mst := structtagmapper.MapOf(reflect.TypeOf(E{}).Field(0).Tag)
	fmt.Println(mst)
}
```
output:

`
map[xml:id json:id comment:test comment]
`
