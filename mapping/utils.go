package mapping

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	isTrue = true
)

// MapAllFieldsStrict applies values from one struct fields to another provided via corresponding arguments, strict types
func MapAllFieldsStrict(from, to interface{}) error {
	return MapAllFields(from, to, &isTrue)
}

// MapAllFields applies values from one struct fields to another provided via corresponding arguments, types checking
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
			}
			continue
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
				case reflect.Float32, reflect.Float64:
					of.SetInt(int64(f.Float()))
				case reflect.Bool:
					of.SetInt(map[bool]int64{false: 0, true: 1}[f.Bool()])
				case reflect.String:
					val, err := strconv.Atoi(f.String())
					if err != nil {
						continue
					}
					of.SetInt(int64(val))
				default:
					return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				switch f.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					of.SetUint(uint64(f.Int()))
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					of.SetUint(f.Uint())
				case reflect.Float32, reflect.Float64:
					of.SetUint(uint64(f.Float()))
				case reflect.Bool:
					of.SetUint(map[bool]uint64{false: 0, true: 1}[f.Bool()])
				case reflect.String:
					val, err := strconv.Atoi(f.String())
					if err != nil {
						continue
					}
					of.SetUint(uint64(val))
				default:
					return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
				}
			case reflect.Float32, reflect.Float64:
				switch f.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					of.SetFloat(float64(f.Int()))
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					of.SetFloat(float64(f.Uint()))
				case reflect.Float32, reflect.Float64:
					of.SetFloat(f.Float())
				case reflect.Bool:
					of.SetFloat(map[bool]float64{false: 0.0, true: 1.0}[f.Bool()])
				case reflect.String:
					val, err := strconv.ParseFloat(f.String(), 64)
					if err != nil {
						continue
					}
					of.SetFloat(val)
				default:
					return errors.New("failed to set field " + e.Type().Field(i).Name + " type " + f.Type().String() + " to " + of.Type().String())
				}


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
					case "bool":
						x := map[bool]int64{false: 0, true: 1}[f.Bool()]
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
					case "bool":
						x := map[bool]uint64{false: 0, true: 1}[f.Bool()]
						of.Set(reflect.ValueOf(&x))
					case "string":
						val, err := strconv.Atoi(f.String())
						if err != nil {
							continue
						}
						x := uint64(val)
						of.Set(reflect.ValueOf(&x))
					}
				case "*float32", "*float64":
					switch f.Type().String() {
					case "int", "int8", "int16", "int32", "int64":
						x := float64(f.Int())
						of.Set(reflect.ValueOf(&x))
					case "uint", "uint8", "uint16", "uint32", "uint64":
						x := float64(f.Uint())
						of.Set(reflect.ValueOf(&x))
					case "float32", "float64":
						x := f.Float()
						of.Set(reflect.ValueOf(&x))
					case "bool":
						x := map[bool]float64{false: 0.0, true: 1.0}[f.Bool()]
						of.Set(reflect.ValueOf(&x))
					case "string":
						val, err := strconv.ParseFloat(f.String(), 64)
						if err != nil {
							continue
						}
						of.Set(reflect.ValueOf(&val))
					}
				case "*bool":
					switch f.Type().String() {
					case "int", "int8", "int16", "int32", "int64":
						x := f.Int() == 0
						of.Set(reflect.ValueOf(&x))
					case "uint", "uint8", "uint16", "uint32", "uint64":
						x := f.Uint() == 0
						of.Set(reflect.ValueOf(&x))
					case "float32", "float64":
						x := f.Float() == float64(0)
						of.Set(reflect.ValueOf(&x))
					case "bool":
						x := f.Bool()
						of.Set(reflect.ValueOf(&x))
					case "string":
						val, err := strconv.ParseBool(f.String())
						if err != nil {
							continue
						}
						of.Set(reflect.ValueOf(&val))
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
