package cmd

import (
	goflag "flag"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

const (
	validResources = `

	* cluster

	`
)

var (
	rootLong = `
	srectl - automate frequent process executed by the sre team,
	focused on reducing toil.
	`

	rootShort = `srectl`
)

type RootCmd struct {
	ClusterName  string
	cobraCommand *cobra.Command
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "srectl",
		Short: rootShort,
		Long:  rootLong,
	},
}

// Execute main function to start command
func Execute() {
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})
	if err := rootCommand.cobraCommand.Execute(); err != nil {
		// exit with errors
	}
}

func init() {

	NewCmdRoot(os.Stdout)
}

func NewCmdRoot(out io.Writer) *cobra.Command {

	cmd := rootCommand.cobraCommand

	// defaultClusterName := os.Getenv("CLUSTER_NAME")
	// cmd.PersistentFlags().StringVarP(&rootCommand.clusterName, "name", "", defaultClusterName, "Name of cluster. Overrides KOPS_CLUSTER_NAME environment variable")

	// create subcommands
	cmd.AddCommand(NewCmdDelete(out))

	return cmd
}

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
