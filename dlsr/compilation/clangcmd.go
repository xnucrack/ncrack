package compilation

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/xnucrack/dlsr/parsing"
	"os"
	"path/filepath"
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
		"baseFileName": func(path string) string {
			return strings.TrimSuffix(path, filepath.Ext(path))
		},
	}).ParseFS(fileTemplates, "*")
	if err != nil {
		return err
	}

	for _, source := range c.Sources {
		source.Path = "dlsr_" + source.Path
		err := func() error {
			base := strings.TrimSuffix(source.Path, filepath.Ext(source.Path))

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
