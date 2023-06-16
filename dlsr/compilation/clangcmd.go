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
	t, err := template.New("").Funcs(template.FuncMap{
		"baseFileName": func(path string) string {
			return strings.TrimSuffix(path, filepath.Ext(path))
		},
	}).ParseFS(fileTemplates, "*")
	if err != nil {
		return err
	}

	t.ExecuteTemplate(os.Stdout, "header", c.Sources[0])
	fmt.Println("=====================================")
	t.ExecuteTemplate(os.Stdout, "main", c.Sources[0])

	return nil
}
