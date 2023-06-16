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

type ReplacementMethod struct {
	Selector string
	Body     string
}

type Source struct {
	Path         string
	CIncludes    []Include
	ObjCIncludes []Include
	TargetClass  string
	Selectors    map[string]ReplacementMethod
}

type Codebase struct {
	// The set of files in the codebase
	Sources []Source
}
