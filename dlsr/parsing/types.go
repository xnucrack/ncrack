package parsing

type IncludeType int

const (
	IncludeTypeSystem IncludeType = iota
	IncludeTypeLocal
)

type Include struct {
	Name string
	Type IncludeType
}

type Source struct {
	Name         string
	CIncludes    []IncludeType
	ObjCIncludes []Include
	TargetClass  string
	Selectors    map[string]string
}

type Codebase struct {
	// The set of files in the codebase
	Sources []Source
}
