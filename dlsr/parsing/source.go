package parsing

import "strings"

func cParseFileDependencies(codeLines []string) []Include {
	var results []Include

	for _, line := range codeLines {
		// check if the line is a #include
		if strings.HasPrefix(line, "#include") {
			// parse the file path...
			filePath := strings.TrimPrefix(line, "#include")
			filePath = strings.TrimSpace(filePath)
			// check if the file path is in the codebase
			if strings.HasPrefix(filePath, "<") {
				// system header
				sysRawFilePath := strings.TrimPrefix(strings.TrimSuffix(filePath, ">"), "<")

				results = append(results, Include{
					Name: sysRawFilePath,
					Type: IncludeTypeSystem,
				})
				continue
			}
			// local header
			filePath = strings.TrimPrefix(filePath, "\"")
			filePath = strings.TrimSuffix(filePath, "\"")

			// add the file to our results...
			results = append(results, Include{
				Name: filePath,
				Type: IncludeTypeLocal,
			})
		}
	}
	return results
}

func objcParseFileDependencies(codeLines []string) []Include {
	var results []Include

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
			sysRawFilePath := strings.TrimPrefix(strings.TrimSuffix(line, ">"), "<")

			results = append(results, Include{
				Name: sysRawFilePath,
				Type: IncludeTypeSystem,
			})
			continue
		}

		// parse the file path...
		filePath := strings.TrimSpace(line)
		filePath = strings.TrimPrefix(filePath, "\"")
		filePath = strings.TrimSuffix(filePath, "\"")
		// add the file to our results...
		results = append(results, Include{
			Name: filePath,
			Type: IncludeTypeLocal,
		})
	}

	return results
}

func ParseSource(path string) (*Source, error) {
	fileExtension := extensionFromName(nameFromPath(path))

	isObjc := stringInSlice(fileExtension, []string{"m", "mm"})

	// read the file contents...
	contents, err := readFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(contents, "\n")

	result := &Source{
		Path: path,
	}

	if isObjc {
		result.ObjCIncludes = objcParseFileDependencies(lines)
	} else {
		result.CIncludes = cParseFileDependencies(lines)
	}

	return result, nil
}
