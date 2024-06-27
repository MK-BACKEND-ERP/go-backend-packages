package maker

import (
	"reflect"
	"strings"
)

// ENuloOuVazio verifica se um valor é nulo ou vazio.
func ENuloOuVazio(valor interface{}) bool {
	if valor == nil {
		return true
	}

	v := reflect.ValueOf(valor)

	switch v.Kind() {
	case reflect.String:
		return strings.TrimSpace(v.String()) == ""
	case reflect.Map, reflect.Slice, reflect.Array:
		return v.Len() == 0
	case reflect.Struct:
		return reflect.DeepEqual(valor, reflect.Zero(reflect.TypeOf(valor)).Interface())
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	default:
		return false
	}
}
