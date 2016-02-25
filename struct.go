package test

type Iface interface {
	Foo()
}

type Struct struct {
	Field1 string
	Field2 int
	Field3 []string
	Field4 uint64
	Field5 string
	Field6 string
	Field7 []byte
}

func (s *Struct) Foo() {}
