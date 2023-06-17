package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/xnucrack/dlsr/compilation"
	"github.com/xnucrack/dlsr/parsing"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run dlsr",
	RunE: func(cmd *cobra.Command, args []string) error {
		includeDir, err := cmd.Flags().GetString("include")
		if err != nil {
			return err
		}

		targetLibrary, err := cmd.Flags().GetString("target")
		if err != nil {
			return err
		}

		if targetLibrary == "" {
			return errors.New("target library cannot be empty")
		}

		links, err := cmd.Flags().GetStringSlice("links")
		if err != nil {
			return err
		}

		frameworks, err := cmd.Flags().GetStringSlice("frameworks")
		if err != nil {
			return err
		}

		// Do the parsing here

		source := parsing.Source{
			Path: "first.m",
			CIncludes: []parsing.Include{
				{"<stdio.h>", parsing.IncludeTypeSystem},
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
				"withName:forAge:": {
					Selector: "withName:forAge:",
					Body: `- (void)withName:(NSString*)name forAge:(NSNumber*)age {
NSLog(@"Name: %@; age: %@", name, age);
}`,
				},
			},
		}
		sources := []parsing.Source{source}

		// Once all the sources are passed, call compilation.Compile with those sources

		return compilation.Compile(parsing.Codebase{
			Sources:       sources,
			IncludePath:   includeDir,
			TargetLibrary: targetLibrary,
			Links:         links,
			Frameworks:    frameworks,
		})
	},
}

func init() {
	runCmd.Flags().StringP("target", "t", "", "target library")
	runCmd.Flags().StringP("include", "i", ".", "includes directory")
	runCmd.Flags().StringSliceP("links", "l", []string{}, "link with the following libraries")
	runCmd.Flags().StringSliceP("frameworks", "f", []string{}, "link with the following frameworks")

	rootCmd.AddCommand(runCmd)
}
