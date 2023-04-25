package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use: "go-studies-grpc",
	}

	cmd.AddCommand(
		Api,
		Client,
	)

	cmd.Execute()
}
