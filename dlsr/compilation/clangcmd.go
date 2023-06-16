package compilation

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/xnucrack/dlsr/parsing"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

//go:embed *.gotxt
var fileTemplates embed.FS

func Compile(c parsing.Codebase) error {
	dir, err := os.MkdirTemp("/tmp/", "dlsr")
	if err != nil {
		return fmt.Errorf("error creating temp directory: %v", err)
	}
	// defer os.RemoveAll(dir)

	t, err := template.New("").Funcs(template.FuncMap{
		"getInterfaceName": func(path string) string {
			baseName := strings.TrimSuffix(path, filepath.Ext(path))
			return "DLSR" + strings.ToTitle(baseName[:1]) + baseName[1:]
		},
		"getMethod": func(body string) string {
			idx := strings.Index(body, "{")
			if idx != -1 {
				return strings.TrimRight(body[:idx], " ") + ";"
			}
			return ""
		},
		"baseFileName": func(path string) string {
			return strings.TrimSuffix(path, filepath.Ext(path))
		},
		"cleanSelector": func(sel string) string {
			return regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(sel, "")
		},
	}).ParseFS(fileTemplates, "*")
	if err != nil {
		return err
	}

	for _, source := range c.Sources {
		source.OutPath = "dlsr_" + source.Path
		err := func() error {
			base := strings.TrimSuffix(source.OutPath, filepath.Ext(source.OutPath))

			headerFile, err := os.Create(filepath.Join(dir, base+".h"))
			if err != nil {
				return err
			}
			defer headerFile.Close()

			t.ExecuteTemplate(headerFile, "header", source)

			mainFile, err := os.Create(filepath.Join(dir, base+".m"))
			if err != nil {
				return err
			}
			defer headerFile.Close()
			t.ExecuteTemplate(mainFile, "main", source)
			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}
