package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Romeoteni188/opsx/internal/system"
	"github.com/spf13/cobra"
)

func newSystemCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "system",
		Short: "Show system information",
		RunE: func(cmd *cobra.Command, args []string) error {
			info := system.Info()

			if jsonOutput {
				output, err := json.MarshalIndent(info, "", "  ")
				if err != nil {
					return err
				}

				fmt.Println(string(output))
				return nil
			}

			fmt.Println("OPSX System")
			fmt.Println()

			fmt.Printf("OS           %s\n", info.OS)
			fmt.Printf("Architecture %s\n", info.Architecture)
			fmt.Printf("Kernel       %s\n", info.Kernel)
			fmt.Printf("CPU          %s\n", info.CPU)
			fmt.Printf("Memory       %s\n", info.Memory)
			fmt.Printf("Disk         %s\n", info.Disk)

			return nil
		},
	}

	cmd.Flags().BoolVar(
		&jsonOutput,
		"json",
		false,
		"Output system information as JSON",
	)

	return cmd
}
