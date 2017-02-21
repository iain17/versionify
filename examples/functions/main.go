package main

import (
	"github.com/Pocketbrain/versionify"
)

type FuncMethod func() string

/**
Called when a method is added for a version, to check a possible set constraint.
*/
func (m FuncMethod) Check(v *versionify.Version) bool {
	return true
}

func main() {
	v := versionify.New()
	v1, _ := v.NewVersion("1.0")
	v2, _ := v.NewVersion("2.0")

	v1.Method("hello", FuncMethod(
		func() string {
			return "my name is Iain"
		},
	))

	v1.Method("goodbye", FuncMethod(
		func() string {
			return "Farewell godspeed."
		},
	))

	v2.Method("hello", FuncMethod(
		func() string {
			return "my name is George"
		},
	))

	methods := v.GetMethods(v2)
	println(methods["hello"].(FuncMethod)())
	println(methods["goodbye"].(FuncMethod)())
}
