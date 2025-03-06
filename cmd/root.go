package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "app",
	Short: "Start social TODO service",
	Run: func(cmd *cobra.Command, args []string) {
		appCtx := newAppContext()
	}
}

func startGinServer(ac appctx.AppContext) {
	ginServer := ac.MustGet()
}
