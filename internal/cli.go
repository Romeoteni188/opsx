package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Name    = "opsx"
	Version = "0.1.0"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     Name,
		Short:   "Linux-first DevOps & DevSecOps CLI",
		Version: Version,
		Long: `OPSX is a Linux-first DevOps and DevSecOps CLI
for everyday infrastructure operations.`,
	}

	rootCmd.AddCommand(newVersionCommand())

	return rootCmd
}

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show OPSX version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("OPSX v%s\n", Version)
		},
	}
}
