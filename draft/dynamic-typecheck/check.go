package dynamictypecheck

import "reflect"

func DynamicTypeCheck() {
	type any interface{}
	var b any = map[string]interface{}{"order": map[string]int{"name": 1}, "age": 1}

	var slice []any
	slice = append(slice, 1, "hoang", struct{ class int }{class: 1}, b)

	for _, element := range slice {
		element := reflect.ValueOf(element)
		element.Type()
		switch element.Kind() {
		case reflect.Map:
			// fmt.Println(element.Type().Key())
			// fmt.Println(element.Type().Elem())
			interable := element.MapRange()
			for interable.Next() {
				if interable.Value().Kind() == reflect.Interface {
				}
			}

		}
	}
}
