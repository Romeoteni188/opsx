package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Romeoteni188/opsx/internal/doctor"
	"github.com/spf13/cobra"
)

func newDoctorCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Check your development environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			report := doctor.Run()

			if jsonOutput {
				output, err := json.MarshalIndent(report, "", "  ")
				if err != nil {
					return err
				}

				fmt.Println(string(output))
				return nil
			}

			fmt.Println("OPSX Doctor")
			fmt.Println()

			for _, category := range report.Categories {
				fmt.Println(category.Name)

				for _, check := range category.Checks {
					if check.Passed {
						fmt.Printf("  ✓ %-20s %s\n", check.Name, check.Message)
					} else {
						fmt.Printf("  ✗ %-20s %s\n", check.Name, check.Message)
					}
				}

				fmt.Println()
			}

			fmt.Printf(
				"Summary\n  %d checks\n  %d passed\n  %d failed\n",
				report.Total(),
				report.Passed(),
				report.Failed(),
			)

			return nil
		},
	}

	cmd.Flags().BoolVar(
		&jsonOutput,
		"json",
		false,
		"Output doctor results as JSON",
	)

	return cmd
}
