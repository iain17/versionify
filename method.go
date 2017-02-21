package versionify

//A method is a functionality available per Version.
//See the methods directory for ready to use methods.
type Method interface {
	Check(v *Version) bool
}
