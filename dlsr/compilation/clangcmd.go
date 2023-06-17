package compilation

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"github.com/xnucrack/dlsr/parsing"
	"os"
	"os/exec"
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

	mainSourceFiles := make([]string, len(c.Sources))

	for i, source := range c.Sources {
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

			mainSourceFiles[i] = base + ".m"

			t.ExecuteTemplate(mainFile, "main", source)
			return nil
		}()
		if err != nil {
			return err
		}
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("cannot change directory: %v", err)
	}

	args := []string{"-dynamiclib", "-I", c.IncludePath}
	if len(c.Frameworks) > 0 {
		for _, framework := range c.Frameworks {
			args = append(args, "-framework")
			args = append(args, framework)
		}
	}
	if len(c.Links) > 0 {
		for _, link := range c.Links {
			args = append(args, "-l")
			args = append(args, link)
		}
	}
	for _, source := range mainSourceFiles {
		args = append(args, source)
	}

	args = append(args, "-o", c.TargetLibrary)

	stderr := new(bytes.Buffer)

	cmd := exec.Command("clang", args...)
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("compilation error: %s", stderr.String())
	}

	_ = currentDir

	return nil
}
