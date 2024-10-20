/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MeMetoCoco3/navi/fvView"

	"github.com/spf13/cobra"
)

// fvCmd represents the fv command
var fvCmd = &cobra.Command{
	Use:   "fv",
	Short: "Show favorites for easy change of directory",
	Long: `Opens the TUI with limited functionality, just allows to delete from favorites with 'a'
  and navigate to the folders you have registered. `,
	Run: func(cmd *cobra.Command, args []string) {
		fvViews.Run()
	},
}

func init() {
	rootCmd.AddCommand(fvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
