package structtagmapper

import (
	"reflect"
	"strconv"
)

// MapOf returns a map of all keys/value pairs in the tag string.
// If the tag does not have the conventional format, the value
// returned by GetMap is unspecified.
// Based on reflect.StructTag.Get(key string)
func MapOf(tag reflect.StructTag) map[string]string {
	result := make(map[string]string, 0)
	for tag != "" {
		// skip leading space
		i := 0
		for i < len(tag) && tag[i] == ' ' {
			i++
		}
		tag = tag[i:]
		if tag == "" {
			break
		}

		// scan to colon.
		// a space or a quote is a syntax error
		i = 0
		for i < len(tag) && tag[i] != ' ' && tag[i] != ':' && tag[i] != '"' {
			i++
		}
		if i+1 >= len(tag) || tag[i] != ':' || tag[i+1] != '"' {
			break
		}
		name := string(tag[:i])
		tag = tag[i+1:]

		// scan quoted string to find value
		i = 1
		for i < len(tag) && tag[i] != '"' {
			if tag[i] == '\\' {
				i++
			}
			i++
		}
		if i >= len(tag) {
			break
		}
		qvalue := string(tag[:i+1])
		tag = tag[i+1:]

		value, _ := strconv.Unquote(qvalue)
		result[name] = value
	}
	return result
}
