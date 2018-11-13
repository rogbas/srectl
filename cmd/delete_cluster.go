package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

type DeleteClusterOptions struct {
	Region      string
	ClusterName string
}

var (
	deleteClusterLong = `
	Delete openshift cluster, removing it's: 
		* Autoscaling Groups
		* Launch Configurations
		* EC2s
		* ELBs
		* VPC
	`

	deleteClusterExample = `
		# Delete a cluster 
		srectl delete cluster my-cluster
	`

	deleteClusterShort = "Delete Cluster"
)

func NewCmdDeleteCluster(out io.Writer) *cobra.Command {
	options := &DeleteClusterOptions{}

	deleteClusterCmd := &cobra.Command{
		Use:     "cluster",
		Short:   deleteClusterShort,
		Long:    deleteClusterLong,
		Example: deleteClusterExample,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) == 0 {
				exitWithError(errors.New("Missing cluster name"))
			}
			options.ClusterName = args[0]

			err := RunDeleteCluster(out, options)
			if err != nil {
				// exitWithError(err)
			}
		},
	}

	deleteClusterCmd.Flags().StringVar(&options.Region, "region", options.Region, "region")

	return deleteClusterCmd
}

func RunDeleteCluster(out io.Writer, options *DeleteClusterOptions) error {

	fmt.Fprintf(out, "Deleting cluster %s on %s", options.ClusterName, options.Region)
	return nil
}
