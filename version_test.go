package versionify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testMethod string

func (m testMethod) Check(v *Version) bool {
	//Test senario. Constraint set on version 3.
	if v.String() == "3.0.0" {
		return false
	}
	return m != "iain"
}

func TestVersion_Method(t *testing.T) {
	//Setup
	v := New()
	v1, err := v.NewVersion("1.0")
	assert.NoError(t, err)

	foo := testMethod("Foo")
	bar := testMethod("Bar")
	iain := testMethod("iain")

	//Call
	_, err = v1.Method("a", foo)
	assert.NoError(t, err)
	_, err = v1.Method("b", bar)
	assert.NoError(t, err)
	_, err = v1.Method("c", iain)
	assert.Error(t, err, "If constraint returns false. Method should not be added.")

	//I should receive both A and B
	methods := v1.GetMethods()
	assert.NotNil(t, methods["a"])
	assert.NotNil(t, methods["b"])
	assert.Nil(t, methods["c"])
	assert.Equal(t, methods["a"], foo)
	assert.Equal(t, methods["b"], bar)
}
