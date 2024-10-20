/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rvCmd represents the rv command
var rvCmd = &cobra.Command{
	Use:   "rv",
	Short: "Remove a directory from favorites.",
	Long:  `If a directory is in favorites, it will remove it. \nIf it is not, it will print a message`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Rv called!")
	},
}

func init() {
	rootCmd.AddCommand(rvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
