package structtagmapper

import (
	"reflect"
	"testing"
)

var tagGetMapTests = []struct {
	Tag reflect.StructTag
	Map map[string]string
}{
	{`protobuf:"PB(1,2)"`, map[string]string{
		"protobuf": "PB(1,2)",
	}},
	{`protobuf:"PB(1,2)" json:"name"`, map[string]string{
		"protobuf": "PB(1,2)",
		"json":     "name",
	}},
	{`   protobuf:"PB(1,2)"      json:"name"     xml:"-"`, map[string]string{
		"protobuf": "PB(1,2)",
		"json":     "name",
		"xml":      "-",
	}},
}

func TestMapOf(t *testing.T) {
	for _, tt := range tagGetMapTests {
		tagmap := MapOf(tt.Tag)
		if !reflect.DeepEqual(tagmap, tt.Map) {
			t.Errorf("MappedStructTag(%#q).GetMap() = %#q, want %#q", tt.Tag, tagmap, tt.Map)
		}

	}
}
