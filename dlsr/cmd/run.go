package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xnucrack/dlsr/compilation"
	"github.com/xnucrack/dlsr/parsing"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run dlsr",
	RunE: func(cmd *cobra.Command, args []string) error {
		source := parsing.Source{
			Path: "first.m",
			CIncludes: []parsing.Include{
				{"<stdio.h>", parsing.IncludeTypeSystem},
				{"\"paddle.h\"", parsing.IncludeTypeLocal},
			},
			ObjCIncludes: []parsing.Include{
				{"<Foundation/Foundation.h>", parsing.IncludeTypeSystem},
			},
			TargetClass: "Cracker",
			Selectors: map[string]parsing.ReplacementMethod{
				"toReplace": {
					Selector: "toReplace",
					Body: `- (void)toReplace {
NSLog(@"I am replaced");
}`,
				},
			},
		}
		sources := []parsing.Source{source}
		return compilation.Compile(parsing.Codebase{Sources: sources})
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
