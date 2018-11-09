package main

import (
	goflag "flag"

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
	focused on reducing the toil.
	`

	rootShort = `srectl`
)

type RootCmd struct {
	clusterName  string
	cobraCommand *cobra.Command
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "srectl",
		Short: rootShort,
		Long:  rootLong,
	},
}

func Execute() {
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})
	if err := rootCommand.cobraCommand.Execute(); err != nil {
		exitWithError(err)
	}
}
