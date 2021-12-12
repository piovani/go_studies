package cmd

import "github.com/spf13/cobra"

var ConsumerWorker = &cobra.Command{
	Use:                   "consumer_worker",
	Short:                 "djadisjasij",
	Version:               "1.0.0",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		initWorker()
	},
}

func initWorker() {

}
