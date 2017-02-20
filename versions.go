package versionify

//This allows the versions to be sorted.
type Versions []*Version

func (v Versions) Len() int {
	return len(v)
}

func (v Versions) Less(i, j int) bool {
	return v[i].LessThan(&v[j].Version)
}

func (v Versions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}