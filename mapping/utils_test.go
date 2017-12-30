package mapping_test

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"mailoman/go-utils/mapping"
)

type case0 struct {
}

type case1 struct {
	Field1 string    `json:"field1" valid:"uuid"`
	Time1  time.Time `json:"time1"`
	Field2 int64     `json:"field2"`
	Field3 string    `json:"field3" valid:"uuid"`
	Field4 bool      `json:"field4"`
}

type case2 struct {
	Field1 *string    `json:"field1" valid:"uuid"`
	Time1  *time.Time `json:"time1"`
	Field2 *int64     `json:"field2"`
	Field3 *string    `json:"field3"`
	Field4 *bool      `json:"field4"`
}

type case3 struct {
	Field1 string     `json:"field1" valid:"uuid"`
	Time1  *time.Time `json:"time1"`
	Field2 *int64     `json:"field2"`
	Field5 *[]string  `json:"field5"`
}

type case4 struct {
	Field1          *string    `json:"field1" valid:"uuid"`
	Time1           *time.Time `json:"time1"`
	Field7          *float32   `json:"field7"`
	Field8          *uint64    `json:"field8"`
	FieldNotInModel *int       `json:"fieldNotInModel"`
}

type case5 struct {
	Field9  uint64 `json:"field9"`
	Field10 int64  `json:"field10"`
	Field11 string `json:"field11"`
	Field12 string `json:"field12"`
}

type case6 struct {
	Field9  float32 `json:"field9"`
	Field10 float64 `json:"field10"`
	Field11 float64 `json:"field11"`
	Field12 float32 `json:"field12"`
}

type case7 struct {
	Field13 uint32 `json:"field13"`
	Field14 uint64 `json:"field14"`
}

type case8 struct {
	Field13 uint64 `json:"field13"`
	Field14 uint32 `json:"field14"`
}

type case9 struct {
	Field13 int  `json:"field13"`
	Field14 int8 `json:"field14"`
}

type case10 struct {
	Field2  *time.Time `json:"field2"`
	Field8  *time.Time `json:"field8"`
	Field13 *time.Time `json:"field13"`
	Field14 *time.Time `json:"field14"`
	Field3  *time.Time `json:"field3"`
	Field4  *time.Time `json:"field4"`
}

type case11 struct {
	Field9  int64  `json:"field9"`
	Field10 uint64 `json:"field10"`
}

type case12 struct {
	Field0  int64  `json:"field0"`
}

type output struct {
	Field0  case0             `json:"field0"`
	Field1  string            `json:"field1" valid:"uuid"`
	Time1   time.Time         `json:"time1"`
	Field2  int64             `json:"field2"`
	Field3  string            `json:"field3" valid:"uuid"`
	Field4  bool              `json:"field4"`
	Field5  []string          `json:"field5"`
	Field6  map[string]string `json:"field6"`
	Field7  float64           `json:"field7"`
	Field8  int32             `json:"field8"`
	Field9  *int64            `json:"field9"`
	Field10 *uint64           `json:"field10"`
	Field11 *int64            `json:"field11"`
	Field12 *uint64           `json:"field12"`
	Field13 uint32            `json:"field13"`
	Field14 uint64            `json:"field14"`
}

var (
	uid  = "83b4b4e2-566c-413e-9be9-98d87e287242"
	now  = time.Now()
	orig = &output{
		Field1: "bfe9c7b3-0c88-4287-881a-02dd2b5c60f7",
		Time1:  now.Add(time.Minute),
		Field2: 1,
		Field3: "value 3",
		Field4: false,
		Field5: []string{"a", "b", "c"},
		Field6: map[string]string{"a1": "a", "b1": "b", "c1": "c"},
		Field7: 10.0,
		Field8: 32,
	}
)

func TestMapAllFieldsStrict_Fail_1(t *testing.T) {
	a := assert.New(t)
	from := case1{}
	to := 1

	err := mapping.MapAllFieldsStrict(from, to)

	a.NotNil(err)
	a.Equal(errors.New("output is not a struct"), err)
}

func TestMapAllFieldsStrict_Fail_2(t *testing.T) {
	a := assert.New(t)
	from := 1
	to := case1{}

	err := mapping.MapAllFieldsStrict(from, to)

	a.NotNil(err)
	a.Equal(errors.New("input is not a struct"), err)
}

func TestMapAllFieldsStrict_Fail_3(t *testing.T) {
	a := assert.New(t)
	from := case1{}
	to := case0{}

	err := mapping.MapAllFieldsStrict(from, to)

	a.NotNil(err)
	a.Equal(errors.New("failed to modify output"), err)
}

func TestMapAllFieldsStrict_Fail_4_1(t *testing.T) {
	a := assert.New(t)
	//subcase1
	from := case10{
		Field2: &now,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.NotNil(err)
	a.Equal(errors.New("failed to set field Field2 type time.Time to int64"), err)
}

func TestMapAllFieldsStrict_Fail_4_2(t *testing.T) {
	a := assert.New(t)
	//subcase2
	from2 := case10{
		Field8: &now,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from2, to)
	a.NotNil(err)
	a.Equal(errors.New("failed to set field Field8 type time.Time to int32"), err)
}

func TestMapAllFieldsStrict_Fail_4_3(t *testing.T) {
	a := assert.New(t)
	//subcase3
	from := case10{
		Field13: &now,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.NotNil(err)
	a.Equal(errors.New("failed to set field Field13 type time.Time to uint32"), err)
}

func TestMapAllFieldsStrict_Fail_5(t *testing.T) {
	a := assert.New(t)
	from := case12{
		Field0: 1,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.NotNil(err)
	a.Equal(errors.New("failed to set field Field0 type int64 to mapping_test.case0"), err)
}

func TestMapAllFieldsStrict_Case1(t *testing.T) {
	// from values to values
	a := assert.New(t)
	from := case1{
		Field1: "83b4b4e2-566c-413e-9be9-98d87e287242",
		Time1:  time.Now(),
		Field2: 2,
		Field3: "new value 3",
		Field4: true,
	}
	to := &output{
		Field1: "bfe9c7b3-0c88-4287-881a-02dd2b5c60f7",
		Time1:  time.Now().Add(time.Minute),
		Field2: 1,
		Field3: "value 3",
		Field4: false,
		Field5: []string{"a", "b", "c"},
		Field6: map[string]string{"a1": "a", "b1": "b", "c1": "c"},
	}
	old := &output{}
	*old = *to

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(from.Field1, to.Field1)
	a.Equal(from.Time1, to.Time1)
	a.Equal(from.Field2, to.Field2)
	a.Equal(from.Field3, to.Field3)
	a.Equal(from.Field4, to.Field4)
	a.Equal(old.Field5, to.Field5)
	a.Equal(old.Field6, to.Field6)
}

func TestMapAllFieldsStrict_Case2(t *testing.T) {
	// from refs to values
	a := assert.New(t)
	uid := "83b4b4e2-566c-413e-9be9-98d87e287242"
	now := time.Now()
	i := int64(2)
	str := "new value 3"
	bool := true
	from := case2{
		Field1: &uid,
		Time1:  &now,
		Field2: &i,
		Field3: &str,
		Field4: &bool,
	}
	to := &output{
		Field1: "bfe9c7b3-0c88-4287-881a-02dd2b5c60f7",
		Time1:  time.Now().Add(time.Minute),
		Field2: 1,
		Field3: "value 3",
		Field4: false,
		Field5: []string{"a", "b", "c"},
		Field6: map[string]string{"a1": "a", "b1": "b", "c1": "c"},
	}
	old := &output{}
	*old = *to

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(*from.Field1, to.Field1)
	a.Equal(*from.Time1, to.Time1)
	a.Equal(*from.Field2, to.Field2)
	a.Equal(*from.Field3, to.Field3)
	a.Equal(*from.Field4, to.Field4)
	a.Equal(old.Field5, to.Field5)
	a.Equal(old.Field6, to.Field6)
}

func TestMapAllFieldsStrict_Cases3_1(t *testing.T) {
	// from values & refs to values
	a := assert.New(t)
	// subcase1
	from := case3{
		Field1: uid,
		Time1:  &now,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(from.Field1, to.Field1)
	a.Equal(*from.Time1, to.Time1)
	a.Equal(orig.Field2, to.Field2)
	a.Equal(orig.Field5, to.Field5)
	a.Equal(orig.Field6, to.Field6)
}

func TestMapAllFieldsStrict_Cases3_2(t *testing.T) {
	// from values & refs to values
	a := assert.New(t)
	i := int64(2)
	// subcase2
	from := case3{
		Field1: uid,
		Field2: &i,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(from.Field1, to.Field1)
	a.Equal(orig.Time1, to.Time1)
	a.Equal(*from.Field2, to.Field2)
	a.Equal(orig.Field5, to.Field5)
	a.Equal(orig.Field6, to.Field6)
}

func TestMapAllFieldsStrict_Cases3_3(t *testing.T) {
	// from values & refs to values
	a := assert.New(t)
	arr := []string{"d", "e"}
	// subcase3
	from := case3{
		Field1: uid,
		Field5: &arr,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(from.Field1, to.Field1)
	a.Equal(orig.Time1, to.Time1)
	a.Equal(orig.Field2, to.Field2)
	a.Equal(*from.Field5, to.Field5)
	a.Equal(orig.Field6, to.Field6)
}

func TestMapAllFieldsStrict_Cases3_4(t *testing.T) {
	// from values & refs to values
	a := assert.New(t)
	// subcase4
	from2 := case4{
		Field1: &uid,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from2, to)
	a.Nil(err)
}

func TestMapAllFieldsStrict_Cases3_5(t *testing.T) {
	// from values & refs with unmatched fields to values
	a := assert.New(t)
	nim := int(123)
	// subcase5
	from2 := case4{
		Field1:          &uid,
		FieldNotInModel: &nim,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from2, to)
	a.NotNil(err)
}

func TestMapAllFieldsStrict_Cases3_6(t *testing.T) {
	// from refs to values & refs, cross-type soft
	a := assert.New(t)
	f := float32(11.0)
	// subcase6
	from2 := case4{
		Field1: &uid,
		Field7: &f,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from2, to)
	a.Nil(err)
	a.Equal(*from2.Field1, to.Field1)
	a.Equal(orig.Time1, to.Time1)
	a.Equal(orig.Field2, to.Field2)
	a.Equal(orig.Field6, to.Field6)
	a.Equal(float64(*from2.Field7), to.Field7)
}

func TestMapAllFieldsStrict_Cases3_7(t *testing.T) {
	// from refs to values & refs, cross-type soft
	a := assert.New(t)
	ui64 := uint64(64)
	// subcase7
	from2 := case4{
		Field1: &uid,
		Field8: &ui64,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from2, to)
	a.Nil(err)
	a.Equal(*from2.Field1, to.Field1)
	a.Equal(orig.Time1, to.Time1)
	a.Equal(orig.Field2, to.Field2)
	a.Equal(orig.Field6, to.Field6)
	a.Equal(int32(*from2.Field8), to.Field8)
}

func TestMapAllFieldsStrict_Cases3_8(t *testing.T) {
	// from refs with unmatched fields to values
	a := assert.New(t)
	nim := int(123)
	// subcase8
	from2 := case4{
		Field1:          &uid,
		FieldNotInModel: &nim,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFields(from2, to, nil)
	a.Nil(err)
	a.Equal(*from2.Field1, to.Field1)
	a.Equal(orig.Time1, to.Time1)
	a.Equal(orig.Field2, to.Field2)
	a.Equal(orig.Field6, to.Field6)
	a.Equal(orig.Field7, to.Field7)
}

func TestMapAllFieldsStrict_Cases4_1(t *testing.T) {
	// from values to refs, cross-type soft
	a := assert.New(t)
	ui64 := uint64(123)
	i64 := int64(123)
	// subcase1
	from := case5{
		Field9:  ui64,
		Field10: i64,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field9, uint64(*to.Field9))
	a.Equal(from.Field10, int64(*to.Field10))
}

func TestMapAllFieldsStrict_Cases4_2(t *testing.T) {
	// from values to refs, cross-type harder
	a := assert.New(t)
	str := "123"
	// subcase2
	from := case5{
		Field11: str,
		Field12: str,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field11, strconv.FormatInt(*to.Field11, 10))
	a.Equal(from.Field12, strconv.FormatInt(int64(*to.Field12), 10))
}

func TestMapAllFieldsStrict_Cases4_3(t *testing.T) {
	// from values to refs, cross-type hard
	a := assert.New(t)
	f32 := float32(123)
	f64 := float64(123)
	// subcase3
	from := case6{
		Field9:  f32,
		Field10: f64,
		Field11: f64,
		Field12: f32,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field9, float32(*to.Field9))
	a.Equal(from.Field10, float64(*to.Field10))
	a.Equal(from.Field11, float64(*to.Field11))
	a.Equal(from.Field12, float32(*to.Field12))
}

func TestMapAllFieldsStrict_Cases5_1(t *testing.T) {
	// from values to values
	a := assert.New(t)
	ui32 := uint32(23)
	ui64 := uint64(46)
	// subcase1
	from := case7{
		Field13: ui32,
		Field14: ui64,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field13, to.Field13)
	a.Equal(from.Field14, to.Field14)
}

func TestMapAllFieldsStrict_Cases5_2(t *testing.T) {
	// from values to values, cross-type soft
	a := assert.New(t)
	ui32 := uint32(23)
	ui64 := uint64(46)
	// subcase2
	from := case8{
		Field13: ui64,
		Field14: ui32,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field13, uint64(to.Field13))
	a.Equal(from.Field14, uint32(to.Field14))
}

func TestMapAllFieldsStrict_Cases5_3(t *testing.T) {
	// from values to values, cross-type soft
	a := assert.New(t)
	i := int(23)
	i8 := int8(46)
	// subcase3
	from := case9{
		Field13: i,
		Field14: i8,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field13, int(to.Field13))
	a.Equal(from.Field14, int8(to.Field14))
}

func TestMapAllFieldsStrict_Cases6_1(t *testing.T) {
	// from values to refs
	a := assert.New(t)
	i64 := int64(23)
	ui64 := uint64(123)
	from := case11{
		Field9:  i64,
		Field10: ui64,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(orig.Field1, to.Field1)
	a.Equal(from.Field9, *to.Field9)
	a.Equal(from.Field10, *to.Field10)
}
