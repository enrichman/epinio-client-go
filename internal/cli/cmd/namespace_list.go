package cmd

import (
	"context"

	"github.com/enrichman/epinio-client-go/internal/cli/ui"
	"github.com/enrichman/epinio-client-go/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type namespaceLister interface {
	List(ctx context.Context) ([]client.Namespace, error)
}

func newNamespaceListCmd(config *Config, namespaceLister namespaceLister) *cobra.Command {
	namespaceListConfig := config.NamespaceConfig.NamespaceListConfig

	namespaceListCmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.List(cmd.OutOrStdout(), namespaceLister)
		},
	}

	bindNamespaceListConfig(namespaceListCmd.Flags(), namespaceListCmd.PersistentFlags(), namespaceListConfig)

	return namespaceListCmd
}

func bindNamespaceListConfig(flags, persistentFlags *pflag.FlagSet, config *NamespaceListConfig) {
	flags.StringVar(&config.Name, "name2", config.Name, "Set your name")
}
