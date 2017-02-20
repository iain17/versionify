package versionify

import (
	"sort"
	"fmt"
)

type VersionRegistrator func(version *Version, methods Methods)

//The entry point of this package. Start a new instance of this and add versions.
type Versionify interface{
	NewVersion(ver string) (*Version, error)
	sort()
	GetMethods(version *Version) Methods
	Register(registrator VersionRegistrator)
}

type versionify struct {
	versions Versions
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
	v.versions = append(v.versions, version)
	v.sort()
	return version, nil
}

/**
Sort the versions. See versions.go to understand how its being sorted.
 */
func (v *versionify) sort() {
	sort.Sort(v.versions)
}

/**
Get the methods of a version. Taking into account previous versions and inheriting the methods of these versions.
 */
func (v *versionify) GetMethods(version *Version) Methods {
	methods := Methods{}
	for _, curVersion := range v.versions {
		if curVersion.GreaterThan(&version.Version) {
			break
		}
		for k, v := range curVersion.GetMethods() {
			if !v.Check(version) {
				fmt.Println("Method did not pass constraint")
				continue
			}
			methods[k] = v
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