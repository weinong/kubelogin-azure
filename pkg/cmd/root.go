package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd provides a cobra root command
func NewRootCmd(version string) *cobra.Command {

	cmd := &cobra.Command{
		Use:          "kubelogin-azure",
		Short:        "login to azure active directory and populate kubeconfig with AAD tokens",
		SilenceUsage: true,
		Version:      version,
		RunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}

	cmd.AddCommand(NewConvertCmd())
	cmd.AddCommand(NewTokenCmd())
	cmd.AddCommand(NewRemoveTokenCacheCmd())

	return cmd
}
