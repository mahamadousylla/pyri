package structs

import (
	"reflect"
)

// GetTags given the tag name to look for, returns all the tags on a struct
// and excludes the tags that are not wanted
func GetTags(src interface{}, tagName string, tagsToIgnore []string) []string {
	m := map[string]bool{}
	for _, s := range tagsToIgnore {
		m[s] = true
	}

	value := reflect.TypeOf(src).Elem()
	return getTags(value, tagName, m)
}

func getTags(t reflect.Type, tagName string, m map[string]bool) []string {
	var result []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)

		if _, ok := m[tag]; ok {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			result = append(result, getTags(field.Type, tagName, m)...)
		}

		if hasTag(tag) {
			result = append(result, tag)
		}
	}

	return result
}

// hasTag if a tag is empty or
// has a '-' it will be ignored
func hasTag(tag string) bool {
	if tag == "" || tag == "-" {
		return false
	}

	return true
}
