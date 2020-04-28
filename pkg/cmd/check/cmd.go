package check

import (
	"github.com/sanchezl/p2pnc/pkg/cmd/check/endpoint"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Runs a check.",
	}
	cmd.AddCommand(endpoint.New())
	return cmd
}
