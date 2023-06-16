package parsing

import "strings"

type Codebase struct {
	// The path to the codebase
	Path string
	// The set of files in the codebase
	Files []string
}

func cParseFileDependencies(codeLines []string) []string {
	var results []string

	for _, line := range codeLines {
		// check if the line is a #include
		if strings.HasPrefix(line, "#include") {
			// parse the file path...
			filePath := strings.TrimPrefix(line, "#include")
			filePath = strings.TrimSpace(filePath)
			// check if the file path is in the codebase
			if strings.HasPrefix(filePath, "<") {
				// system header
				results = append(results, filePath)
				continue
			}
			// local header
			filePath = strings.TrimPrefix(filePath, "\"")
			filePath = strings.TrimSuffix(filePath, "\"")
			// add the file to our results...
			results = append(results, filePath)
		}
	}
	return results
}

func objcParseFileDependencies(codeLines []string) []string {
	var results []string

	for _, line := range codeLines {
		// check if the line is an #import

		if !strings.HasPrefix(line, "@") && !strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "import") {
			line = strings.TrimPrefix(line, "import")
		} else if strings.HasPrefix(line, "include") {
			line = strings.TrimPrefix(line, "include")
		} else {
			continue
		}

		// check if it's a system header...
		if strings.HasPrefix(line, "<") {
			// system header
			results = append(results, line)
			continue
		}

		// parse the file path...
		filePath := strings.TrimSpace(line)
		filePath = strings.TrimPrefix(filePath, "\"")
		filePath = strings.TrimSuffix(filePath, "\"")
		// add the file to our results...
		results = append(results, filePath)
	}

	return results

}

func (c *Codebase) addFileInternal(path string) {
	// check if the file is already in our list...
	for _, file := range c.Files {
		if file == path {
			return
		}
	}
	// append to our list...
	c.Files = append(c.Files, path)
}

func ParseCodebase(path string) {
	// read in
}
