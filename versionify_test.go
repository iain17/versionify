package versionify

import (
	"gopkg.in/stretchr/testify.v1/assert"
	"testing"
)

func TestVersionify_GetMethods(t *testing.T) {
	//Setup
	v := New()
	v1, err := v.NewVersion("1.0")
	assert.NoError(t, err)
	v2, err := v.NewVersion("2.0")
	assert.NoError(t, err)
	v3, err := v.NewVersion("3.0")
	assert.NoError(t, err)
	v4, err := v.NewVersion("4.0")
	assert.NoError(t, err)

	AV1 := testMethod("A on V1")
	BV1 := testMethod("B on V1")
	BV2 := testMethod("B on V2")
	CV2 := testMethod("C on V2")

	_, err = v1.Method("a", &AV1)
	assert.NoError(t, err)
	_, err = v1.Method("b", &BV1)
	assert.NoError(t, err)

	_, err = v2.Method("b", &BV2)
	assert.NoError(t, err)
	_, err = v2.Method("c", &CV2)
	assert.NoError(t, err)

	//Call
	methods := v.GetMethods(v2)

	assert.NotNil(t, methods["a"], "Should have inherited method a from v1")
	assert.Equal(t, methods["b"], &BV2, "Should have overwritten method b from v1 with v2")
	assert.NotNil(t, methods["c"], "Should have its own method c")

	methodsV3 := v.GetMethods(v3)
	assert.Equal(t, 0, len(methodsV3), "Due to the testMethod. No methods are on v3.")

	methodsV4 := v.GetMethods(v4)
	assert.Equal(t, 3, len(methodsV4), "Due to the testMethod. Eventhough we haven't added any. all methods are inherited")
}

func TestVersionify_NewVersion(t *testing.T) {
	v := New()
	_, err := v.NewVersion("1.0")
	assert.NoError(t, err)
	_, err = v.NewVersion("2.0")
	assert.NoError(t, err)
	_, err = v.NewVersion("2.0")
	assert.Error(t, err, "Version 2 is already registered!")
	_, err = v.NewVersion("abc")
	assert.Error(t, err, "This is not a real version?")

}

func TestVersionify_Register(t *testing.T) {
	//Setup
	v := New()
	v2, err := v.NewVersion("2.0")
	assert.NoError(t, err)
	v1, err := v.NewVersion("1.0")
	assert.NoError(t, err)

	v1Called := false
	v2Called := false

	v.Register(func(version *Version, methods Methods) {
		if !v1Called {
			assert.Equal(t, version, v1, "First version should be called.")
			assert.Equal(t, methods, v.GetMethods(v1))
			v1Called = true
			return
		}

		if !v2Called {
			assert.Equal(t, version, v2, "Second version should be called after v1.")
			assert.Equal(t, methods, v.GetMethods(v2))
			v2Called = true
			return
		}
	})

	assert.True(t, v1Called, "V1 should have been registered")
	assert.True(t, v2Called, "V2 should have been registered")
}
