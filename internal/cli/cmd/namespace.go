package cmd

import (
	"github.com/spf13/cobra"
)

type namespaceService interface {
	namespaceLister
}

func newNamespaceCmd(config *Config, namespaceService namespaceService) *cobra.Command {
	namespaceCmd := &cobra.Command{
		Use:   "namespace",
		Short: "A brief description of your application",
	}

	namespaceCmd.AddCommand(newNamespaceListCmd(config, namespaceService))

	return namespaceCmd
}
