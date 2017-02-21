package string_versionify

import (
	"github.com/Pocketbrain/versionify"
)

//The most simple example where we make a primitive now a usable versionify method.
//This is done just by implementing the Check method.
type StringMethod string

/**
Called when a method is added for a version, to check a possible set constraint.
*/
func (m *StringMethod) Check(v *versionify.Version) bool {
	return true
}
