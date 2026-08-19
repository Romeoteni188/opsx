package cli

import "github.com/spf13/cobra"

const Version = "0.2.0"

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show OPSX version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("OPSX v" + Version)
		},
	}
}
