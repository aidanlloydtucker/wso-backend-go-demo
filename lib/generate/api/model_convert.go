package api

import (
	"reflect"
	"strings"
)

func ModelToSchema(model interface{}, schema interface{}) reflect.Value {
	modelType := reflect.TypeOf(model)
	schemaType := reflect.TypeOf(schema)

	m := reflect.ValueOf(model)
	var s reflect.Value


	if modelType.Kind() == reflect.Slice && schemaType.Kind() == reflect.Slice {
		s = reflect.MakeSlice(reflect.SliceOf(schemaType.Elem()), m.Len(), m.Cap())

		for i := 0; i < s.Len(); i++ {
			s.Index(i).Set(ModelToSchema(m.Index(i).Interface(), s.Index(i).Interface()))
		}
	} else if modelType.Kind() == reflect.Struct && schemaType.Kind() == reflect.Struct {
		schemaPtr := reflect.New(schemaType)
		s = schemaPtr.Elem()

		for i := 0; i < schemaType.NumField(); i++ {
			field := schemaType.Field(i)
			modelFieldLoc := field.Tag.Get("api")

			var val reflect.Value
			if modelFieldLoc != "" {
				val = getFieldByLoc(strings.Split(modelFieldLoc, "."), m)
			} else {
				val = m.FieldByName(field.Name)
			}

			s.FieldByName(field.Name).Set(val)
		}

	}

	return s
}

func getFieldByLoc(loc []string, val reflect.Value) reflect.Value {
	field := val.FieldByName(loc[0])
	if len(loc) == 1 {
		return field
	}

	return getFieldByLoc(loc[1:], val)
}