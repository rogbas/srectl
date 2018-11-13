package cmd

import (
	"io"

	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	deleteLong = `
	Delete instance
	`

	deleteExample = `
		# Delete a cluster 
		srectl delete cluster my-cluster
	`

	deleteShort = "Delete function"
)

type DeleteOptions struct {
}

func NewCmdDelete(out io.Writer) *cobra.Command {
	options := &DeleteOptions{}

	cmd := &cobra.Command{
		Use:        "delete",
		Short:      deleteShort,
		Long:       deleteLong,
		Example:    deleteExample,
		SuggestFor: []string{"rm"},
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(RunDelete(out, options))
		},
	}

	// Add subcommands
	cmd.AddCommand(NewCmdDeleteCluster(out))

	return cmd
}

func RunDelete(out io.Writer, d *DeleteOptions) error {

	return nil
}
