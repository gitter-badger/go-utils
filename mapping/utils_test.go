package mapping_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mailoman/go-utils/mapping"
)

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
	Field8          *uint64      `json:"field8"`
	FieldNotInModel *int       `json:"fieldNotInModel"`
}


type output struct {
	Field1 string            `json:"field1" valid:"uuid"`
	Time1  time.Time         `json:"time1"`
	Field2 int64             `json:"field2"`
	Field3 string            `json:"field3" valid:"uuid"`
	Field4 bool              `json:"field4"`
	Field5 []string          `json:"field5"`
	Field6 map[string]string `json:"field6"`
	Field7 float64           `json:"field7"`
	Field8 int32             `json:"field8"`
}

var (
	uid = "83b4b4e2-566c-413e-9be9-98d87e287242"
	now = time.Now()
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

func TestMapAllFieldsStrict_Case1(t *testing.T) {
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
	a := assert.New(t)

	from := case3{
		Field1: uid,
		Time1:  &now,
	}
	to := &output{}
	*to = *orig

	// subcase1
	err := mapping.MapAllFieldsStrict(from, to)
	a.Nil(err)
	a.Equal(from.Field1, to.Field1)
	a.Equal(*from.Time1, to.Time1)
	a.Equal(orig.Field2, to.Field2)
	a.Equal(orig.Field5, to.Field5)
	a.Equal(orig.Field6, to.Field6)
}

func TestMapAllFieldsStrict_Cases3_2(t *testing.T) {
	a := assert.New(t)
	// subcase2
	i := int64(2)
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
	*to = *orig
}

func TestMapAllFieldsStrict_Cases3_3(t *testing.T) {
	a := assert.New(t)
	// subcase3
	arr := []string{"d", "e"}
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
	a := assert.New(t)
	// subcase5
	nim := int(123)
	from2 := case4{
		Field1: &uid,
		//Field8: &u,
		FieldNotInModel: &nim,
	}
	to := &output{}
	*to = *orig

	err := mapping.MapAllFieldsStrict(from2, to)
	a.NotNil(err)
}

func TestMapAllFieldsStrict_Cases3_6(t *testing.T) {
	a := assert.New(t)
	// subcase6
	f := float32(11.0)
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
	//a.Equal(int64(*from2.Field8), to.Field8)
}

func TestMapAllFieldsStrict_Cases3_7(t *testing.T) {
	a := assert.New(t)
	// subcase7
	i64 := uint64(64)
	from2 := case4{
		Field1: &uid,
		Field8: &i64,
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
	a := assert.New(t)
	// subcase8
	nim := int(123)
	from2 := case4{
		Field1: &uid,
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
