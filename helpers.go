package godi

import (
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

// returns the number as a string and its size, which will never be more than 19 characters (int64)
func isNumericValue(refValue *reflect.Value) (numStr string, size uint8, ok bool) {

	switch refValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		numStr = strconv.FormatInt(refValue.Int(), 10)
		ok = true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		numStr = strconv.FormatUint(refValue.Uint(), 10)
		ok = true

	case reflect.Float32:
		buf := make([]byte, 0, 10) // allocate a 32-byte buffer to avoid reallocation
		buf = strconv.AppendFloat(buf, refValue.Float(), 'f', -1, 32)
		numStr = unsafe.String(unsafe.SliceData(buf), len(buf))
		ok = true
	case reflect.Float64:
		buf := make([]byte, 0, 19) // allocate a 64-byte buffer to avoid reallocation
		buf = strconv.AppendFloat(buf, refValue.Float(), 'f', -1, 64)
		numStr = unsafe.String(unsafe.SliceData(buf), len(buf))
		ok = true

	}

	return numStr, uint8(len(numStr)), ok
}

func snakeCase(str string) string {
	var result string
	for i, r := range str {
		if unicode.IsUpper(r) {
			// Si es mayúscula y no es el primer caracter, añadimos "_"
			if i > 0 && (unicode.IsLower(rune(str[i-1])) || unicode.IsDigit(rune(str[i-1]))) {
				result += "_"
			}
			// Convertimos la mayúscula a minúscula
			result += strings.ToLower(string(r))
		} else {
			// Si no es mayúscula, simplemente lo añadimos
			result += string(r)
		}
	}
	return result
}
