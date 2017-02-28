package versionify

import (
	"errors"
	"sort"
	"fmt"
)

type VersionRegistrator func(version *Version, methods Methods)

//The entry point of this package. Start a new instance of this and add versions.
type Versionify interface {
	NewVersion(ver string) (*Version, error)
	sort()
	exists(ver *Version) bool
	GetMethods(version *Version) Methods
	Register(registrator VersionRegistrator)
	SetReverse(value bool)
}

type versionify struct {
	versions Versions
	reverse bool
}

func New() Versionify {
	return &versionify{
		versions: Versions{},
	}
}

/**
Register a new version. Returns a Version object which can then be used to register methods on.
*/
func (v *versionify) NewVersion(ver string) (*Version, error) {
	version, err := newVersion(ver)
	if err != nil {
		return nil, err
	}
	if v.exists(version) {
		return nil, errors.New("This version is already registered.")
	}

	v.versions = append(v.versions, version)
	v.sort()
	return version, nil
}

/**
Check if version isn't already registered.
*/
func (v *versionify) exists(ver *Version) bool {
	for _, version := range v.versions {
		if version.Equal(&ver.Version) {
			return true
		}
	}
	return false
}

/**
Sort the versions. See versions.go to understand how its being sorted.
*/
func (v *versionify) sort() {
	if v.reverse {
		sort.Sort(sort.Reverse(v.versions))
	} else {
		sort.Sort(v.versions)
	}
}

/**
Reverse sorting. This will make the
 */
func (v *versionify) SetReverse(value bool) {
	v.reverse = value
}

/**
Get the methods of a version. Taking into account previous versions and inheriting the methods of these versions.
*/
func (v *versionify) GetMethods(version *Version) Methods {
	methods := Methods{}
	for _, curVersion := range v.versions {
		//Fail safe. Should not happen due to the sorting
		if v.reverse {
			if curVersion.LessThan(&version.Version) {
				continue
			}
		} else {
			if curVersion.GreaterThan(&version.Version) {
				continue
			}
		}

		for k, method := range curVersion.GetMethods() {
			if !method.Check(version) {
				fmt.Printf("Method '%s' from v%s did not pass constraint for v%s\n", k, curVersion.String(), version.String())
				continue
			}
			methods[k] = method
		}
	}
	return methods
}

//This function will go through each version and call the registrator.
//This allows us to easily register methods.
func (v *versionify) Register(registrator VersionRegistrator) {
	for _, version := range v.versions {
		registrator(version, v.GetMethods(version))
	}
}
