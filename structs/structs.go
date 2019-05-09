package structs

import (
	"reflect"

	fstructs "github.com/fatih/structs"
)

// GetTags given the tag name to look for, returns all the tags on a struct
// and excludes the tags that are not wanted
func GetTags(src interface{}, tagName string, tagsToIgnore []string) []string {
	d := map[string]bool{}
	for _, s := range tagsToIgnore {
		d[s] = true
	}

	t := reflect.TypeOf(src).Elem()
	return getTags(t, tagName, d)
}

// GetTagsAndValues: given the tag name to look for,
// the tags the exclude and the name of all nested structs,
// returns all the tags and values on a struct
func GetTagsAndValues(src interface{}, tagName string, tagsToIgnore, nestedStructNames []string) ([]string, []interface{}) {
	d := map[string]bool{}
	for _, s := range tagsToIgnore {
		d[s] = true
	}

	n := map[string]bool{}
	for _, s := range nestedStructNames {
		n[s] = true
	}

	t := reflect.TypeOf(src).Elem()

	tags := getTags(t, tagName, d)
	values := getValues(src, tagName, d, n)

	return tags, values
}

func getTags(t reflect.Type, tagName string, d map[string]bool) []string {
	var result []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(tagName)

		if _, ok := d[tag]; ok {
			continue
		}

		if field.Type.Kind() == reflect.Struct {
			result = append(result, getTags(field.Type, tagName, d)...)
		}

		if hasTag(tag) {
			result = append(result, tag)
		}
	}

	return result
}

func getValues(src interface{}, tagName string, d map[string]bool, n map[string]bool) []interface{} {
	var result []interface{}

	fields := fstructs.Fields(src)

	for _, field := range fields {
		tag := field.Tag(tagName)
		if _, ok := d[tag]; ok {
			continue
		}

		k := field.Value()

		if _, ok := n[reflect.TypeOf(k).Name()]; ok {
			result = append(result, getValues(k, tagName, d, n)...)
		} else {
			result = append(result, k)
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
