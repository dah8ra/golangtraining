package ex122

import (
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x), 0, 0)
}


var bound int = 15
var loopbound int = 17

// var numbound int = 3

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value, count int, loopcount int) {
	if loopcount > loopbound {
		fmt.Println("Skip!!!")
		return
	}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), count, loopcount)
		}
	case reflect.Struct:
		////////////////
		// Added loop count check
		////////////////
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			loopcount++
			if loopcount < loopbound {
				display(fieldPath, v.Field(i), count, loopcount)
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			///////////////////////////
			// Added a case of array about map key.
			///////////////////////////
			// fmt.Println("@@@ ", key.Type())
			// fmt.Println("@@@ ", key.Kind().String())
			keytype := key.Kind().String()
			if "array" == keytype {
				for i := 0; i < key.Len(); i++ {
					display(fmt.Sprintf("%s[%d]", path, i), v.MapIndex(key), count, loopcount)
				}
			} else {
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key), count, loopcount)
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), count, loopcount)
		}

	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), count, loopcount)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
		//////////////
		// Added if clause
		//////////////
		/*
			if !strings.Contains(path, "uncommon") {
				fmt.Printf("%s = %s\n", path, formatAtom(v))
			}
		*/
	}
}
