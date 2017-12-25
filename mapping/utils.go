package mapping

import (
	"errors"
	"reflect"
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
				return errors.New("failed to find field in output: " + e.Type().Field(i).Name)
			} else {
				continue
			}
		}

		of := outs.FieldByName(e.Type().Field(i).Name)

		if of.Kind() == reflect.Ptr && of.IsValid() && of.CanSet() {// && !of.IsNil()
			of = of.Elem()
		}

		if of.IsValid() && of.CanSet() {
			switch of.Kind() {
			case reflect.String:
				of.SetString(f.String())
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
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
			case reflect.Bool:
				of.SetBool(f.Bool())
			case reflect.Map, reflect.Array, reflect.Slice, reflect.Interface, reflect.Struct:
				of.Set(f)
			default:
				of.Set(f)
				//return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
			}
		} else {
			if strict != nil && *strict == true {
				return errors.New("failed to find field in output: " + e.Type().Field(i).Name)
			}
			continue
		}
	}
	return nil
}
