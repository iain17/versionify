package versionify

//A method is a functionality available per Version.
type Method interface {
	Check(v *Version) bool
}
