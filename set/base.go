package set


type Set interface {

	Add(items... interface{}) bool

	Remove(items... interface{}) bool
}