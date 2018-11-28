package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws/session"
	awsasg "github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/rogbas/srectl/pkg/cloud/aws/autoscaling"
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

	fmt.Printf("Deleting cluster %s on %s\n", options.ClusterName, options.Region)
	// instances := []*awsec2.Instance{}
	// What needs to be deleted:
	// 	* AutoScaling Groups
	//  * Launch Configs
	// 	* Master EC2s
	//	* ELBs
	// 	* VPC

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// ec2Svc := ec2.NewService(awsec2.New(sess))

	asgSvc := autoscaling.NewService(awsasg.New(sess))

	err := asgSvc.DeleteAutoScalingGroupsByCluster(options.ClusterName)
	if err != nil {
		fmt.Println("Error deleting AutoScaling Groups", err)
	}

	return nil
}
