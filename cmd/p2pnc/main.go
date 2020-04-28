package main

import (
	"github.com/sanchezl/p2pnc/pkg/cmd/check"
	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()
	cmd := new()
	cmd.Execute()
}

func new() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "p2pnc",
		Short: "Point-to-point network check tool",
	}
	cmd.AddCommand(check.New())
	return cmd
}
