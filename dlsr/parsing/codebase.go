package parsing

func ParseCodebase(path string) (*Codebase, error) {
	// scan the directory for code files...
	result := &Codebase{}

	sourceFileExensions := []string{"c", "h", "m", "mm"}
	sourceFiles := filesInDirectory(path, &sourceFileExensions)

	for _, sourceFile := range sourceFiles {
		parsed, err := ParseSource(sourceFile)
		if err != nil {
			continue
		}

		result.Sources = append(result.Sources, *parsed)
	}

	return result, nil
}
