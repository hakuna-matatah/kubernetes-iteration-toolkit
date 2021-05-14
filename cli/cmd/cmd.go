package main

import (
	"github.com/prateekgogia/kit/cli/pkg/util/logging"
	"github.com/spf13/cobra"
)

var log = logging.NewNamedLogger("KIT")

// cmd represents the base command when called without any subcommands
var cmd = &cobra.Command{
	Use:   "kit-cli",
	Short: "Command line interface to the Kubernetes Iteration Toolkit",
	Long:  ``,
}

func init() {
	cmd.AddCommand(
		Apply,
		Delete,
	)
}

func main() {
	cobra.CheckErr(cmd.Execute())
}
