package cli

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "opsx",
		Short:   "Linux-first DevOps & DevSecOps CLI",
		Version: Version,
		Long: `OPSX is a Linux-first DevOps and DevSecOps CLI
for everyday infrastructure operations.`,
	}

	rootCmd.AddCommand(
		newVersionCommand(),
		newDoctorCommand(),
		newSystemCommand(),
	)

	return rootCmd
}
