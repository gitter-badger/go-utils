package mapping

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	isTrue = true
)

func MapAllFieldsStrict(from, to interface{}) error {
	return MapAllFields(from, to, &isTrue)
}

func MapAllFields(from, to interface{}, strict *bool) error {
	e := reflect.ValueOf(from)
	outs := reflect.ValueOf(to)

	if outs.Kind() == reflect.Ptr {
		outs = outs.Elem()
	}

	if outs.Kind() != reflect.Struct {
		return errors.New("output is not a struct")
	}
	if e.Kind() != reflect.Struct {
		return errors.New("input is not a struct")
	}

	if !outs.CanSet() {
		return errors.New("failed to modify output")
	}

	ofields := map[string]reflect.StructField{}
	for i := 0; i < outs.NumField(); i++ {
		ofields[outs.Type().Field(i).Name] = outs.Type().Field(i)
	}

	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)

		if !f.IsValid() {
			continue
		}

		if f.Kind() == reflect.Ptr {
			if !f.IsNil() {
				f = f.Elem()
			} else {
				continue
			}
		}

		if _, ok := ofields[e.Type().Field(i).Name]; !ok {
			if strict != nil && *strict == true {
				return errors.New("failed to find field in output[1]: " + e.Type().Field(i).Name)
			} else {
				continue
			}
		}

		of := outs.FieldByName(e.Type().Field(i).Name)

		if of.CanSet() {
			switch of.Kind() {
			case reflect.String:
				of.SetString(f.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uintptr:
				switch f.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					of.SetInt(f.Int())
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					of.SetInt(int64(f.Uint()))
				default:
					return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				switch f.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					of.SetUint(uint64(f.Int()))
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					of.SetUint(f.Uint())
				default:
					return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
				}
			case reflect.Float32, reflect.Float64:
				of.SetFloat(f.Float())
			case reflect.Ptr:
				switch of.Type().String() {
				case "*int", "*int8", "*int16", "*int32", "*int64":
					switch f.Type().String() {
					case "int", "int8", "int16", "int32", "int64":
						x := int64(f.Int())
						of.Set(reflect.ValueOf(&x))
					case "uint", "uint8", "uint16", "uint32", "uint64":
						x := int64(f.Uint())
						of.Set(reflect.ValueOf(&x))
					case "float32", "float64":
						x := int64(f.Float())
						of.Set(reflect.ValueOf(&x))
					case "string":
						val, err := strconv.Atoi(f.String())
						if err != nil {
							continue
						}
						x := int64(val)
						of.Set(reflect.ValueOf(&x))
					}
				case "*uint", "*uint8", "*uint16", "*uint32", "*uint64":
					switch f.Type().String() {
					case "int", "int8", "int16", "int32", "int64":
						x := uint64(f.Int())
						of.Set(reflect.ValueOf(&x))
					case "uint", "uint8", "uint16", "uint32", "uint64":
						x := uint64(f.Uint())
						of.Set(reflect.ValueOf(&x))
					case "float32", "float64":
						x := uint64(f.Float())
						of.Set(reflect.ValueOf(&x))
					case "string":
						val, err := strconv.Atoi(f.String())
						if err != nil {
							continue
						}
						x := uint64(val)
						of.Set(reflect.ValueOf(&x))
					}
				}
			case reflect.Bool:
				of.SetBool(f.Bool())
			case reflect.Map, reflect.Array, reflect.Slice, reflect.Interface, reflect.Struct:
				if f.Type() != of.Type() {
					return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
				}
				of.Set(f)
			default:
				of.Set(f)
				//return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
			}
		} else {
			if strict != nil && *strict == true {
				return errors.New("failed to find field in output[2]: " + e.Type().Field(i).Name)
			}
			continue
		}
	}
	return nil
}
