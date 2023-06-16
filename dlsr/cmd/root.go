package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "dlsr",
	Short: "Dylib Symbol Replacer",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	return rootCmd.Execute()
}
