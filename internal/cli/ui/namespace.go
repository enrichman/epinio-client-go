package ui

import (
	"context"
	"fmt"
	"io"

	"github.com/enrichman/epinio-client-go/pkg/client"
)

type namespaceLister interface {
	List(ctx context.Context) ([]client.Namespace, error)
}

func List(writer io.Writer, namespaceLister namespaceLister) error {
	namespaces, err := namespaceLister.List(context.Background())
	if err != nil {
		return err
	}

	fmt.Fprintln(writer, "namespaceList: ", namespaces)
	return nil
}
