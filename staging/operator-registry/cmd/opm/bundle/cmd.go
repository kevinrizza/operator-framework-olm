package bundle

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Hidden: true,
		Use:    "bundle",
		Short:  "Operator bundle commands",
		Long:   `Generate operator bundle metadata and build bundle image.`,
	}

	runCmd.AddCommand(newBundleGenerateCmd())
	runCmd.AddCommand(newBundleBuildCmd())
	return runCmd
}
