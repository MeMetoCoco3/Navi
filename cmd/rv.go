/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/MeMetoCoco3/navi/favorites"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// rvCmd represents the rv command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a directory from favorites.",
	Long:  `If a directory is in favorites, it will remove it. \nIf it is not, it will print a message`,
	Run: func(cmd *cobra.Command, args []string) {
		cd, err := os.Getwd()
		if err != nil {
			log.Fatalln("(-) Error getting current directory: ", err)
		}

		favorites.RemoveFav(cd)
		log.Fatalf("(+) Directory '%s' was successfully removed", cd)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
