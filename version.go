package versionify

import (
	hversion "github.com/hashicorp/go-version"
	"errors"
)

//This extends hversion.Version so that with each version we can have methods.
type Methods map[string]Method

type Version struct {
	hversion.Version
	methods Methods
}

/**
Used internally by versionify to create a new version.
 */
func newVersion(ver string) (*Version, error) {
	hver, err := hversion.NewVersion(ver)
	if err != nil {
		return nil, err
	}

	return &Version{
		Version: *hver,
		methods: Methods{},
	}, nil
}

/**
Registers a method for this version.
Returns an error if the constraint of that method does not conform the version we are adding it to.
 */
func (v *Version) Method(name string, method Method) (Method, error) {
	if !method.Check(v) {
		return nil, errors.New("Tried adding a method on a version that directly conflicts with the constraint of that method. Edit the constraints and try again.")
	}
	v.methods[name] = method
	return v.methods[name], nil
}

/**
Returns all the methods of this version.
 */
func (v *Version) GetMethods() Methods {
	return v.methods
}