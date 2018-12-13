package internal

type Assertion struct {
	input  interface{}
	offset int
	extra  []interface{}
}

func New(input interface{}, offset int, extra ...interface{}) *Assertion {
	return &Assertion{
		input:  input,
		offset: offset,
		extra:  extra,
	}
}
