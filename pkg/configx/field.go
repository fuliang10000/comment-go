package configx

import (
	"reflect"
	"strings"
)

type fieldInfo struct {
	defaultValues map[string]string
}

func getFieldName(field reflect.StructField) string {
	name := field.Tag.Get(configTagKey)
	if pos := strings.IndexByte(name, tagSepKey); pos >= 0 {
		return strings.ToLower(strings.TrimSpace(name[:pos]))
	}
	return strings.ToLower(strings.TrimSpace(name))
}
func getDefault(field reflect.StructField) *string {
	name := field.Tag.Get(configTagKey)
	if pos := strings.IndexByte(name, tagSepKey); pos >= 0 {
		s1, s2, found := strings.Cut(name[pos:], "=")
		if found && strings.HasSuffix(s1, defaultTagKey) {
			return &s2
		}
	}
	return nil
}

func fullName(field reflect.StructField, name string) string {
	if name == "" {
		return getFieldName(field)
	}
	return name + nameSepKey + getFieldName(field)
}

func (fi *fieldInfo) build(tp reflect.Type, name string) {
	tp = Deref(tp)
	switch tp.Kind() {
	case reflect.Struct:
		fi.buildStruct(tp, name)
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Func:
		return
	default:
		return
	}
}

func (fi *fieldInfo) buildStruct(tp reflect.Type, name string) {
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		if !field.IsExported() {
			continue
		}

		ft := Deref(field.Type)
		// flatten anonymous fields
		if field.Anonymous {
			switch ft.Kind() {
			case reflect.Struct:
				fi.buildStruct(ft, name)
			default:
			}
		} else {
			if getDefault(field) != nil {
				fi.defaultValues[fullName(field, name)] = *getDefault(field)
			}
			fi.build(ft, fullName(field, name))
		}
	}
}

// Deref dereferences a type, if pointer type, returns its element type.
func Deref(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func buildFieldInfo(tp reflect.Type) *fieldInfo {
	var r = &fieldInfo{
		defaultValues: map[string]string{},
	}
	r.build(tp, "")
	return r
}
