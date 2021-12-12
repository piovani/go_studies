package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var API = &cobra.Command{
	Use:                   "api",
	Short:                 "djasi",
	Version:               "1.0.0",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		initService()
	},
}

func initService() {
	fmt.Println("Iniciar API")
}
