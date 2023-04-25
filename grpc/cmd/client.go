package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Client = &cobra.Command{
	Use: "client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLIENT")
	},
}

