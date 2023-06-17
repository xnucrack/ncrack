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
	OutPath      string
	CIncludes    []Include
	ObjCIncludes []Include
	TargetClass  string
	Selectors    map[string]ReplacementMethod
}

type Codebase struct {
	Sources       []Source // The set of sources in the codebase
	IncludePath   string   // Add the following to clang -I path
	TargetLibrary string   // The patch will ocurr on this library
	Links         []string // Link with the following libraries
	Frameworks    []string // Link with the following frameworks
}
