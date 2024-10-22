/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MeMetoCoco3/navi/views"
	"github.com/spf13/cobra"
)

// gatorCmd represents the gator command
var gatorCmd = &cobra.Command{
	Use:   "gator",
	Short: "Navigate your file system",
	Long: `Navigate your file system with Gator command with the arrows 
  of your keyboard. Use 'a' to add folders to favorite list for easy access.
 
  Controls: 
    - Use arrows to move in folders or back.
    - 'a' to add/quit path from favorites.
    - 'g' or 'enter' to CD to selected path.
    - 'q' or 'ctrl+c' to quit.`,
	Run: func(cmd *cobra.Command, args []string) {
		views.Run()
	},
}

func init() {

	rootCmd.AddCommand(gatorCmd)

}
