package ddd

type Scalarable interface {
	ToString() string
	ToScalar() interface{}
}
