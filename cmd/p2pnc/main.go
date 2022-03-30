package main

import (
	"github.com/sanchezl/p2pnc/pkg/cmd/check"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()
	cmd := newP2PNCCommand()
	err := cmd.Execute()
	if err != nil {
		klog.Fatal(err)
	}
}

func newP2PNCCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "p2pnc",
		Short: "Point-to-point network check tool",
	}
	cmd.AddCommand(check.New())
	return cmd
}
