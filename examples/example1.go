package examples

// Input structure
type InputExample1 struct {
	Str string
	I32 int32
	I64 int64
	Boo bool
	F32 float32
}

// Output structure
type OutputExample1 struct {
	Str float64
	I32 string
	I64 string
	Boo int
	F32 int64
}
