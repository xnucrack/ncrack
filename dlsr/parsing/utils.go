package parsing

import (
	"os"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func readFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func nameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func extensionFromName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

// filesInDirectory returns a list of files in the given directory.
func filesInDirectory(path string, extensions *[]string) []string {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var results []string
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		if extensions != nil {
			// split the file name by the dot...
			parts := strings.Split(entry.Name(), ".")
			if len(parts) < 2 {
				continue
			}
			if stringInSlice(parts[len(parts)-1], *extensions) {
				results = append(results, entry.Name())
			} else {
				continue
			}
		}

		results = append(results, entry.Name())
	}
	return results
}
